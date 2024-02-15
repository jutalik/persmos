package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"permos/x/permos/types"
)

func (k msgServer) CreatePermos(goCtx context.Context, msg *types.MsgCreatePermos) (*types.MsgCreatePermosResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var permos = types.Permos{
		Creator:  msg.Creator,
		Mypermos: msg.Mypermos,
	}

	id := k.AppendPermos(
		ctx,
		permos,
	)

	return &types.MsgCreatePermosResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdatePermos(goCtx context.Context, msg *types.MsgUpdatePermos) (*types.MsgUpdatePermosResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var permos = types.Permos{
		Creator:  msg.Creator,
		Id:       msg.Id,
		Mypermos: msg.Mypermos,
	}

	// Checks that the element exists
	val, found := k.GetPermos(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetPermos(ctx, permos)

	return &types.MsgUpdatePermosResponse{}, nil
}

func (k msgServer) DeletePermos(goCtx context.Context, msg *types.MsgDeletePermos) (*types.MsgDeletePermosResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetPermos(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePermos(ctx, msg.Id)

	return &types.MsgDeletePermosResponse{}, nil
}
