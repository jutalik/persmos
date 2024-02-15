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

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PermosKeeper(t)
	permos.InitGenesis(ctx, k, genesisState)
	got := permos.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
