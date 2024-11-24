package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "loan/testutil/keeper"
	"loan/x/load/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.LoadKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
