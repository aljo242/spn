package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) TriggerLaunch(goCtx context.Context, msg *types.MsgTriggerLaunch) (*types.MsgTriggerLaunchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.GetChain(ctx, msg.LaunchID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrChainNotFound, "%d", msg.LaunchID)
	}

	// Check sender is the coordinator of the chain
	coordinatorID, found := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if !found {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordAddressNotFound, msg.Coordinator)
	}
	if chain.CoordinatorID != coordinatorID {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordInvalid, fmt.Sprintf(
			"coordinator of the chain is %d",
			chain.CoordinatorID,
		))
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrapf(types.ErrTriggeredLaunch, "%d", msg.LaunchID)
	}

	if msg.RemainingTime < k.MinLaunchTime(ctx) {
		return nil, sdkerrors.Wrapf(types.ErrLaunchTimeTooLow, "%d", msg.RemainingTime)
	}
	if msg.RemainingTime > k.MaxLaunchTime(ctx) {
		return nil, sdkerrors.Wrapf(types.ErrLaunchTimeTooHigh, "%d", msg.RemainingTime)
	}

	chain.LaunchTriggered = true
	chain.LaunchTimestamp = ctx.BlockTime().Unix() + int64(msg.RemainingTime)
	k.SetChain(ctx, chain)

	return &types.MsgTriggerLaunchResponse{}, nil
}
