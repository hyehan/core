package market

import (
	"testing"

	core "github.com/classic-terra/core/v3/types"
	"github.com/classic-terra/core/v3/x/market/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var randomPrice = sdk.NewDec(1700)

func setup(t *testing.T) (keeper.TestInput, sdk.Handler) {
	input := keeper.CreateTestInput(t)

	params := input.MarketKeeper.GetParams(input.Ctx)
	input.MarketKeeper.SetParams(input.Ctx, params)
	input.OracleKeeper.SetLunaExchangeRate(input.Ctx, core.MicroSDRDenom, randomPrice)
	input.OracleKeeper.SetLunaExchangeRate(input.Ctx, core.MicroKRWDenom, randomPrice)
	h := NewHandler(input.MarketKeeper)

	return input, h
}
