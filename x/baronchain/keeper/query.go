package keeper

import (
	"baronchain/x/baronchain/types"
)

var _ types.QueryServer = Keeper{}
