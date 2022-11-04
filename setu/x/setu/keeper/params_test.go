package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "setu/testutil/keeper"
	"setu/x/setu/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SetuKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
