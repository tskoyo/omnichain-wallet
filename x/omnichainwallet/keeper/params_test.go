package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/tskoyo/omnichain-wallet/testutil/keeper"
	"github.com/tskoyo/omnichain-wallet/x/omnichainwallet/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.OmnichainwalletKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
