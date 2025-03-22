package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "baronchain/testutil/keeper"
	"baronchain/x/evm/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EvmKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
