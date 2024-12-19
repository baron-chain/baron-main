package keeper

import (
	"baronchain/x/tokenbank/types"
)

var _ types.QueryServer = Keeper{}
