package keeper

import (
	"baronchain/x/bcevm/types"
)

var _ types.QueryServer = Keeper{}
