package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"permos/x/permos/types"
)

func (k Keeper) PermosAll(ctx context.Context, req *types.QueryAllPermosRequest) (*types.QueryAllPermosResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var permoss []types.Permos

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	permosStore := prefix.NewStore(store, types.KeyPrefix(types.PermosKey))

	pageRes, err := query.Paginate(permosStore, req.Pagination, func(key []byte, value []byte) error {
		var permos types.Permos
		if err := k.cdc.Unmarshal(value, &permos); err != nil {
			return err
		}

		permoss = append(permoss, permos)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPermosResponse{Permos: permoss, Pagination: pageRes}, nil
}

func (k Keeper) Permos(ctx context.Context, req *types.QueryGetPermosRequest) (*types.QueryGetPermosResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	permos, found := k.GetPermos(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetPermosResponse{Permos: permos}, nil
}
