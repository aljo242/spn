package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgSettleRequest(t *testing.T) {
	var (
		coordinator1                = sample.Coordinator(sample.Address())
		coordinator2                = sample.Coordinator(sample.Address())
		invalidChain                = uint64(1000)
		k, pk, _, srv, _, _, sdkCtx = setupMsgServer(t)
		ctx                         = sdk.WrapSDKContext(sdkCtx)
	)

	coordinator1.CoordinatorID = pk.AppendCoordinator(sdkCtx, coordinator1)
	coordinator2.CoordinatorID = pk.AppendCoordinator(sdkCtx, coordinator2)

	chains := createNChainForCoordinator(k, sdkCtx, coordinator1.CoordinatorID, 3)
	chains[0].LaunchTriggered = true
	k.SetChain(sdkCtx, chains[0])
	chains[1].CoordinatorID = 99999
	k.SetChain(sdkCtx, chains[1])

	requestSamples := make([]RequestSample, 6)
	for i := 0; i < 6; i++ {
		addr := sample.Address()
		requestSamples[i] = RequestSample{
			Content: sample.GenesisAccountContent(chains[2].LaunchID, addr),
			Creator: addr,
		}
	}
	requests := createRequestsFromSamples(k, sdkCtx, chains[2].LaunchID, requestSamples)

	tests := []struct {
		name      string
		msg       types.MsgSettleRequest
		checkAddr string
		err       error
	}{
		{
			name: "invalid chain",
			msg: types.MsgSettleRequest{
				LaunchID:  invalidChain,
				Signer:    coordinator1.Address,
				RequestID: requests[0].RequestID,
				Approve:   true,
			},
			err: types.ErrChainNotFound,
		},
		{
			name: "launch triggered chain",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[0].LaunchID,
				Signer:    coordinator1.Address,
				RequestID: requests[0].RequestID,
				Approve:   true,
			},
			err: types.ErrTriggeredLaunch,
		},
		{
			name: "coordinator not found",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[1].LaunchID,
				Signer:    coordinator1.Address,
				RequestID: requests[0].RequestID,
				Approve:   true,
			},
			err: types.ErrChainInactive,
		},
		{
			name: "no permission error",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Signer:    coordinator2.Address,
				RequestID: requests[0].RequestID,
				Approve:   true,
			},
			err: types.ErrNoAddressPermission,
		},
		{
			name: "should prevent approving an invalid request",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Signer:    coordinator1.Address,
				RequestID: 99999999,
				Approve:   true,
			},
			err: types.ErrRequestNotFound,
		},
		{
			name: "coordinator can approve a request",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Signer:    coordinator1.Address,
				RequestID: requests[0].RequestID,
				Approve:   true,
			},
			checkAddr: requestSamples[0].Creator,
		},
		{
			name: "coordinator can approve a second request for the same chain",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Signer:    coordinator1.Address,
				RequestID: requests[1].RequestID,
				Approve:   true,
			},
			checkAddr: requestSamples[1].Creator,
		},
		{
			name: "coordinator can reject a request",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Signer:    coordinator1.Address,
				RequestID: requests[2].RequestID,
				Approve:   false,
			},
			checkAddr: requestSamples[2].Creator,
		},
		{
			name: "request creator can reject their own request",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Signer:    requestSamples[3].Creator,
				RequestID: requests[3].RequestID,
				Approve:   false,
			},
			checkAddr: requestSamples[3].Creator,
		},
		{
			name: "should prevent rejecting a request if the signer is not the request creator",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Signer:    requestSamples[3].Creator,
				RequestID: requests[4].RequestID,
				Approve:   false,
			},
			err: types.ErrNoAddressPermission,
		},
		{
			name: "request creator if not the coordinator, should not be able to approve their own chain request 8",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Signer:    requestSamples[5].Creator,
				RequestID: requests[5].RequestID,
				Approve:   true,
			},
			err: types.ErrNoAddressPermission,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := srv.SettleRequest(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}
			require.NoError(t, err)

			_, found := k.GetRequest(sdkCtx, tt.msg.LaunchID, tt.msg.RequestID)
			require.False(t, found, "request not removed")

			_, found = k.GetGenesisAccount(sdkCtx, tt.msg.LaunchID, tt.checkAddr)
			require.Equal(t, tt.msg.Approve, found, "request apply not performed")
		})
	}
}
