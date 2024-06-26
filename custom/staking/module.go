package staking

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/staking/types"

	customtypes "github.com/classic-terra/core/v3/custom/staking/types"
	core "github.com/classic-terra/core/v3/types"
)

var _ module.AppModuleBasic = AppModuleBasic{}

// AppModuleBasic defines the basic application module used by the staking module.
type AppModuleBasic struct {
	staking.AppModuleBasic
}

// RegisterLegacyAminoCodec registers the staking module's types for the given codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	customtypes.RegisterLegacyAminoCodec(cdc)
	*types.ModuleCdc = *customtypes.ModuleCdc
}

// DefaultGenesis returns default genesis state as raw bytes for the gov
// module.
func (am AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	// customize to set default genesis state deposit denom to uluna
	defaultGenesisState := types.DefaultGenesisState()
	defaultGenesisState.Params.BondDenom = core.MicroLunaDenom

	return cdc.MustMarshalJSON(defaultGenesisState)
}
