package cosmocheckersibc_test

import (
	"testing"

	keepertest "github.com/malakaja/cosmo-checkers-ibc/testutil/keeper"
	"github.com/malakaja/cosmo-checkers-ibc/testutil/nullify"
	"github.com/malakaja/cosmo-checkers-ibc/x/cosmocheckersibc"
	"github.com/malakaja/cosmo-checkers-ibc/x/cosmocheckersibc/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmocheckersibcKeeper(t)
	cosmocheckersibc.InitGenesis(ctx, *k, genesisState)
	got := cosmocheckersibc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
