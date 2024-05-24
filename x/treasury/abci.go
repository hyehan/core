package treasury

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	core "github.com/classic-terra/core/v3/types"
	"github.com/classic-terra/core/v3/x/treasury/keeper"
	"github.com/classic-terra/core/v3/x/treasury/types"
)

// EndBlocker is called at the end of every block
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	// Burn all coins from the burn module account
	k.BurnCoinsFromBurnAccount(ctx)

	// Check epoch last block
	if !core.IsPeriodLastBlock(ctx, core.BlocksPerWeek) {
		return
	}

	// Update luna issuance after finish all works
	defer k.RecordEpochInitialIssuance(ctx)

	// Compute & Update internal indicators for the current epoch
	k.UpdateIndicators(ctx)

	// Check probation period
	if ctx.BlockHeight() < int64(core.BlocksPerWeek*k.WindowProbation(ctx)) {
		return
	}

	// Settle seigniorage to oracle & distribution(community-pool) module-account
	// TODO: need to consider whether to remove this commenting when the swap enables
	// k.SettleSeigniorage(ctx)

	// Update tax-rate and reward-weight of next epoch
	taxRate := k.UpdateTaxPolicy(ctx)
	rewardWeight := k.UpdateRewardPolicy(ctx)
	taxCap := k.UpdateTaxCap(ctx)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypePolicyUpdate,
			sdk.NewAttribute(types.AttributeKeyTaxRate, taxRate.String()),
			sdk.NewAttribute(types.AttributeKeyRewardWeight, rewardWeight.String()),
			sdk.NewAttribute(types.AttributeKeyTaxCap, taxCap.String()),
		),
	)
}
