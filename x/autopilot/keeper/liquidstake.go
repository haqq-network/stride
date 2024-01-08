package keeper

import (
	"errors"
	"fmt"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"

	"github.com/Stride-Labs/stride/v16/x/autopilot/types"
	stakeibckeeper "github.com/Stride-Labs/stride/v16/x/stakeibc/keeper"
	stakeibctypes "github.com/Stride-Labs/stride/v16/x/stakeibc/types"
)

func (k Keeper) TryLiquidStaking(
	ctx sdk.Context,
	packet channeltypes.Packet,
	newData transfertypes.FungibleTokenPacketData,
	packetMetadata types.StakeibcPacketMetadata,
) error {
	params := k.GetParams(ctx)
	if !params.StakeibcActive {
		return errorsmod.Wrapf(types.ErrPacketForwardingInactive, "autopilot stakeibc routing is inactive")
	}

	// In this case, we can't process a liquid staking transaction, because we're dealing with STRD tokens
	if transfertypes.ReceiverChainIsSource(packet.GetSourcePort(), packet.GetSourceChannel(), newData.Denom) {
		return errors.New("the native token is not supported for liquid staking")
	}

	amount, ok := sdk.NewIntFromString(newData.Amount)
	if !ok {
		return errors.New("not a parsable amount field")
	}

	// Note: newData.denom is base denom e.g. uatom - not ibc/xxx
	var token = sdk.NewCoin(newData.Denom, amount)

	prefixedDenom := transfertypes.GetPrefixedDenom(packet.GetDestPort(), packet.GetDestChannel(), newData.Denom)
	ibcDenom := transfertypes.ParseDenomTrace(prefixedDenom).IBCDenom()

	hostZone, err := k.stakeibcKeeper.GetHostZoneFromHostDenom(ctx, token.Denom)
	if err != nil {
		return err
	}

	if hostZone.IbcDenom != ibcDenom {
		return fmt.Errorf("ibc denom %s is not equal to host zone ibc denom %s", ibcDenom, hostZone.IbcDenom)
	}

	strideAddress, err := sdk.AccAddressFromBech32(packetMetadata.StrideAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid stride_address (%s) in autopilot memo", strideAddress)
	}

	return k.RunLiquidStake(ctx, strideAddress, token, packetMetadata)
}

func (k Keeper) RunLiquidStake(ctx sdk.Context, addr sdk.AccAddress, token sdk.Coin, packetMetadata types.StakeibcPacketMetadata) error {
	msg := &stakeibctypes.MsgLiquidStake{
		Creator:   addr.String(),
		Amount:    token.Amount,
		HostDenom: token.Denom,
	}

	if err := msg.ValidateBasic(); err != nil {
		return err
	}

	msgServer := stakeibckeeper.NewMsgServerImpl(k.stakeibcKeeper)
	result, err := msgServer.LiquidStake(
		sdk.WrapSDKContext(ctx),
		msg,
	)
	if err != nil {
		return errorsmod.Wrapf(err, "failed to liquid stake")
	}

	if packetMetadata.IbcReceiver == "" {
		return nil
	}

	hostZone, err := k.stakeibcKeeper.GetHostZoneFromHostDenom(ctx, token.Denom)
	if err != nil {
		return err
	}

	return k.IBCTransferStAsset(ctx, result.StToken, addr.String(), hostZone, packetMetadata)
}

func (k Keeper) IBCTransferStAsset(ctx sdk.Context, stAsset sdk.Coin, sender string, hostZone *stakeibctypes.HostZone, packetMetadata types.StakeibcPacketMetadata) error {
	ibcTransferTimeoutNanos := k.stakeibcKeeper.GetParam(ctx, stakeibctypes.KeyIBCTransferTimeoutNanos)
	timeoutTimestamp := uint64(ctx.BlockTime().UnixNano()) + ibcTransferTimeoutNanos
	channelId := packetMetadata.TransferChannel
	if channelId == "" {
		channelId = hostZone.TransferChannelId
	}
	transferMsg := &transfertypes.MsgTransfer{
		SourcePort:       transfertypes.PortID,
		SourceChannel:    channelId,
		Token:            stAsset,
		Sender:           sender,
		Receiver:         packetMetadata.IbcReceiver,
		TimeoutTimestamp: timeoutTimestamp,
		// TimeoutHeight and Memo are unused
	}

	_, err := k.transferKeeper.Transfer(sdk.WrapSDKContext(ctx), transferMsg)
	return err
}
