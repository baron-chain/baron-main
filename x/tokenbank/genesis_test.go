package tokenbank_test

import (
	"testing"

	keepertest "baronchain/testutil/keeper"
	"baronchain/testutil/nullify"
	"baronchain/x/tokenbank"
	"baronchain/x/tokenbank/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TokenbankKeeper(t)
	tokenbank.InitGenesis(ctx, *k, genesisState)
	got := tokenbank.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
