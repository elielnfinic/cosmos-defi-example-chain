package load_test

import (
	"testing"

	keepertest "loan/testutil/keeper"
	"loan/testutil/nullify"
	load "loan/x/load/module"
	"loan/x/load/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LoadKeeper(t)
	load.InitGenesis(ctx, k, genesisState)
	got := load.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
