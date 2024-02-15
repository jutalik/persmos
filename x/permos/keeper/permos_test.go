package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "permos/testutil/keeper"
	"permos/testutil/nullify"
	"permos/x/permos/keeper"
	"permos/x/permos/types"
)

func createNPermos(keeper keeper.Keeper, ctx context.Context, n int) []types.Permos {
	items := make([]types.Permos, n)
	for i := range items {
		items[i].Id = keeper.AppendPermos(ctx, items[i])
	}
	return items
}

func TestPermosGet(t *testing.T) {
	keeper, ctx := keepertest.PermosKeeper(t)
	items := createNPermos(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetPermos(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestPermosRemove(t *testing.T) {
	keeper, ctx := keepertest.PermosKeeper(t)
	items := createNPermos(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePermos(ctx, item.Id)
		_, found := keeper.GetPermos(ctx, item.Id)
		require.False(t, found)
	}
}

func TestPermosGetAll(t *testing.T) {
	keeper, ctx := keepertest.PermosKeeper(t)
	items := createNPermos(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPermos(ctx)),
	)
}

func TestPermosCount(t *testing.T) {
	keeper, ctx := keepertest.PermosKeeper(t)
	items := createNPermos(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetPermosCount(ctx))
}
