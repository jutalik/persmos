package permos_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "permos/testutil/keeper"
	"permos/testutil/nullify"
	"permos/x/permos/module"
	"permos/x/permos/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PermosList: []types.Permos{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PermosCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PermosKeeper(t)
	permos.InitGenesis(ctx, k, genesisState)
	got := permos.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PermosList, got.PermosList)
	require.Equal(t, genesisState.PermosCount, got.PermosCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
