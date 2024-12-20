// x/pqc/types/params.go
package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// Parameter keys
var (
	KeyMaxMessageSize = []byte("MaxMessageSize")
)

// Default parameter values
var (
	DefaultMaxMessageSize = uint64(1024) // 1KB default max message size
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMaxMessageSize, &p.MaxMessageSize, validateMaxMessageSize),
	}
}

func validateMaxMessageSize(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max message size cannot be 0")
	}

	return nil
}

// DefaultParams returns default parameters
func DefaultParams() Params {
	return Params{
		MaxMessageSize: DefaultMaxMessageSize,
	}
}
