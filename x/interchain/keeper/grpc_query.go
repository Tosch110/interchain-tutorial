package keeper

import (
	"github.com/tosch110/interchain/x/interchain/types"
)

var _ types.QueryServer = Keeper{}
