package omnichainwallet_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/tskoyo/omnichain-wallet/testutil/keeper"
	"github.com/tskoyo/omnichain-wallet/testutil/nullify"
	omnichainwallet "github.com/tskoyo/omnichain-wallet/x/omnichainwallet/module"
	"github.com/tskoyo/omnichain-wallet/x/omnichainwallet/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OmnichainwalletKeeper(t)
	omnichainwallet.InitGenesis(ctx, k, genesisState)
	got := omnichainwallet.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
