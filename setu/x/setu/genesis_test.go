package setu_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "setu/testutil/keeper"
	"setu/testutil/nullify"
	"setu/x/setu"
	"setu/x/setu/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SetuKeeper(t)
	setu.InitGenesis(ctx, *k, genesisState)
	got := setu.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
