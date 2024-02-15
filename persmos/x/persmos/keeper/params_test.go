package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "persmos/testutil/keeper"
	"persmos/x/persmos/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.PersmosKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
