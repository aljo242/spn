package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringp/keeper"
	"github.com/tendermint/spn/x/monitoringp/types"
)

func createTestConsumerClientID(keeper *keeper.Keeper, ctx sdk.Context) types.ConsumerClientID {
	item := types.ConsumerClientID{}
	keeper.SetConsumerClientID(ctx, item)
	return item
}

func TestConsumerClientIDGet(t *testing.T) {
	keeper, _, ctx := keepertest.MonitoringpKeeper(t)
	item := createTestConsumerClientID(keeper, ctx)
	rst, found := keeper.GetConsumerClientID(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestConsumerClientIDRemove(t *testing.T) {
	keeper, _, ctx := keepertest.MonitoringpKeeper(t)
	createTestConsumerClientID(keeper, ctx)
	keeper.RemoveConsumerClientID(ctx)
	_, found := keeper.GetConsumerClientID(ctx)
	require.False(t, found)
}
