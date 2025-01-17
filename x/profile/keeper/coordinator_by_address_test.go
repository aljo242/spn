package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

func createNCoordinatorByAddress(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.CoordinatorByAddress {
	items := make([]types.CoordinatorByAddress, n)
	for i := range items {
		items[i].Address = sample.Address()
		keeper.SetCoordinatorByAddress(ctx, items[i])
	}
	return items
}

func TestCoordinatorByAddressGet(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	items := createNCoordinatorByAddress(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCoordinatorByAddress(ctx, item.Address)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestCoordinatorByAddressRemove(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	items := createNCoordinatorByAddress(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCoordinatorByAddress(ctx, item.Address)
		_, found := keeper.GetCoordinatorByAddress(ctx, item.Address)
		require.False(t, found)
	}
}

func TestCoordinatorByAddressGetAll(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	items := createNCoordinatorByAddress(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllCoordinatorByAddress(ctx))
}

func TestCoordinatorIDFromAddress(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	address := sample.Address()
	keeper.SetCoordinatorByAddress(ctx, types.CoordinatorByAddress{
		Address:       address,
		CoordinatorID: 10,
	})

	id, found := keeper.CoordinatorIDFromAddress(ctx, address)
	require.True(t, found)
	require.Equal(t, uint64(10), id)

	_, found = keeper.CoordinatorIDFromAddress(ctx, sample.Address())
	require.False(t, found)
}
