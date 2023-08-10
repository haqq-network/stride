package v13

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	stakeibckeeper "github.com/Stride-Labs/stride/v13/x/stakeibc/keeper"
	stakeibctypes "github.com/Stride-Labs/stride/v13/x/stakeibc/types"
)

var (
	UpgradeName = "v13"
)

// CreateUpgradeHandler creates an SDK upgrade handler for v13
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	stakeibcKeeper stakeibckeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info("Starting upgrade v13...")

		ctx.Logger().Info("Registering stTokens to consumer reward denom whitelist...")
		hostZones := stakeibcKeeper.GetAllHostZone(ctx)
		allDenoms := []string{}

		// get all stToken denoms
		for _, zone := range hostZones {
			allDenoms = append(allDenoms, stakeibctypes.StAssetDenomFromHostZoneDenom(zone.HostDenom))
		}

		err := stakeibcKeeper.RegisterStTokenDenomsToWhitelist(ctx, allDenoms)
		if err != nil {
			return nil, errorsmod.Wrapf(err, "unable to register stTokens to whitelist")
		}

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
