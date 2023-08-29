package keeper

import (
	"github.com/malakaja/cosmo-checkers-ibc/x/cosmocheckersibc/types"
)

var _ types.QueryServer = Keeper{}
