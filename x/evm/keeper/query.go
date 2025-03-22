package keeper

import (
	"baronchain/x/evm/types"
)

var _ types.QueryServer = Keeper{}
