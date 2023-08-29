package keeper_test

import (
	"testing"

	testkeeper "github.com/malakaja/cosmo-checkers-ibc/testutil/keeper"
	"github.com/malakaja/cosmo-checkers-ibc/x/cosmocheckersibc/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CosmocheckersibcKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
