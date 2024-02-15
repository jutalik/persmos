package persmos_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "persmos/testutil/keeper"
	"persmos/testutil/nullify"
	"persmos/x/persmos/module"
	"persmos/x/persmos/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PersmosKeeper(t)
	persmos.InitGenesis(ctx, k, genesisState)
	got := persmos.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
