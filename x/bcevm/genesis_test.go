package bcevm_test

import (
	"testing"

	keepertest "baronchain/testutil/keeper"
	"baronchain/testutil/nullify"
	"baronchain/x/bcevm"
	"baronchain/x/bcevm/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BcevmKeeper(t)
	bcevm.InitGenesis(ctx, *k, genesisState)
	got := bcevm.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
