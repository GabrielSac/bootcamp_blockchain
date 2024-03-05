package keeper_test

import (
	"context"
	"testing"

	keepertest "lambchain/testutil/keeper"
	"lambchain/testutil/nullify"
	"lambchain/x/lambchain/keeper"
	"lambchain/x/lambchain/types"

	"github.com/stretchr/testify/require"
)

func createNSuma(keeper keeper.Keeper, ctx context.Context, n int) []types.Suma {
	items := make([]types.Suma, n)
	for i := range items {
		items[i].Id = keeper.AppendSuma(ctx, items[i])
	}
	return items
}

func TestSumaGet(t *testing.T) {
	keeper, ctx := keepertest.LambchainKeeper(t)
	items := createNSuma(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSuma(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSumaRemove(t *testing.T) {
	keeper, ctx := keepertest.LambchainKeeper(t)
	items := createNSuma(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSuma(ctx, item.Id)
		_, found := keeper.GetSuma(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSumaGetAll(t *testing.T) {
	keeper, ctx := keepertest.LambchainKeeper(t)
	items := createNSuma(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSuma(ctx)),
	)
}

func TestSumaCount(t *testing.T) {
	keeper, ctx := keepertest.LambchainKeeper(t)
	items := createNSuma(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSumaCount(ctx))
}
