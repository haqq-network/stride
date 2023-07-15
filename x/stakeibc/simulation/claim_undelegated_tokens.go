package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/Stride-Labs/stride/v12/x/stakeibc/keeper"
	"github.com/Stride-Labs/stride/v12/x/stakeibc/types"
)

func SimulateMsgClaimUndelegatedTokens(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgClaimUndelegatedTokens{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ClaimUndelegatedTokens simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ClaimUndelegatedTokens simulation not implemented"), nil, nil
	}
}
