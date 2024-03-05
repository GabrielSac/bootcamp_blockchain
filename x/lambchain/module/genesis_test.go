package lambchain_test

import (
	"testing"

	keepertest "lambchain/testutil/keeper"
	"lambchain/testutil/nullify"
	lambchain "lambchain/x/lambchain/module"
	"lambchain/x/lambchain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SumaList: []types.Suma{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		SumaCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LambchainKeeper(t)
	lambchain.InitGenesis(ctx, k, genesisState)
	got := lambchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.SumaList, got.SumaList)
	require.Equal(t, genesisState.SumaCount, got.SumaCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
