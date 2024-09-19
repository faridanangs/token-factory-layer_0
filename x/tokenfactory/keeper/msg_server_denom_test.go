package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "tokenfactory/testutil/keeper"
	"tokenfactory/x/tokenfactory/keeper"
	"tokenfactory/x/tokenfactory/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestDenomMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.TokenfactoryKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateDenom{Creator: creator,
			Denom: strconv.Itoa(i),
		}
		_, err := srv.CreateDenom(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetDenom(ctx,
			expected.Denom,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestDenomMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateDenom
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateDenom{Creator: creator,
				Denom: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateDenom{Creator: "B",
				Denom: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateDenom{Creator: creator,
				Denom: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.TokenfactoryKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateDenom{Creator: creator,
				Denom: strconv.Itoa(0),
			}
			_, err := srv.CreateDenom(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateDenom(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetDenom(ctx,
					expected.Denom,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestDenomMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteDenom
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteDenom{Creator: creator,
				Denom: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteDenom{Creator: "B",
				Denom: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteDenom{Creator: creator,
				Denom: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.TokenfactoryKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateDenom(ctx, &types.MsgCreateDenom{Creator: creator,
				Denom: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteDenom(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetDenom(ctx,
					tc.request.Denom,
				)
				require.False(t, found)
			}
		})
	}
}
