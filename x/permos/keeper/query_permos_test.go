package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "permos/testutil/keeper"
	"permos/testutil/nullify"
	"permos/x/permos/types"
)

func TestPermosQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PermosKeeper(t)
	msgs := createNPermos(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetPermosRequest
		response *types.QueryGetPermosResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetPermosRequest{Id: msgs[0].Id},
			response: &types.QueryGetPermosResponse{Permos: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetPermosRequest{Id: msgs[1].Id},
			response: &types.QueryGetPermosResponse{Permos: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetPermosRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Permos(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestPermosQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PermosKeeper(t)
	msgs := createNPermos(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPermosRequest {
		return &types.QueryAllPermosRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PermosAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Permos), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Permos),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PermosAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Permos), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Permos),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.PermosAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Permos),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.PermosAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
