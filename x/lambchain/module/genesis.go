package lambchain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"lambchain/x/lambchain/keeper"
	"lambchain/x/lambchain/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the suma
	for _, elem := range genState.SumaList {
		k.SetSuma(ctx, elem)
	}

	// Set suma count
	k.SetSumaCount(ctx, genState.SumaCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.SumaList = k.GetAllSuma(ctx)
	genesis.SumaCount = k.GetSumaCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
