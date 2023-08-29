package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/malakaja/cosmo-checkers-ibc/testutil/keeper"
	"github.com/malakaja/cosmo-checkers-ibc/x/cosmocheckersibc/keeper"
	"github.com/malakaja/cosmo-checkers-ibc/x/cosmocheckersibc/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmocheckersibcKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
