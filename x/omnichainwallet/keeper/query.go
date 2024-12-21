package keeper

import (
	"github.com/tskoyo/omnichain-wallet/x/omnichainwallet/types"
)

var _ types.QueryServer = Keeper{}
