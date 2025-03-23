package keeper_test

import (
	"testing"

	testkeeper "baronchain/testutil/keeper"
	"baronchain/x/bcevm/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BcevmKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
