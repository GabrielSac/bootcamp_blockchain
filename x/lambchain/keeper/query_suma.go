package keeper

import (
	"context"

	"lambchain/x/lambchain/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SumaAll(ctx context.Context, req *types.QueryAllSumaRequest) (*types.QueryAllSumaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sumas []types.Suma

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	sumaStore := prefix.NewStore(store, types.KeyPrefix(types.SumaKey))

	pageRes, err := query.Paginate(sumaStore, req.Pagination, func(key []byte, value []byte) error {
		var suma types.Suma
		if err := k.cdc.Unmarshal(value, &suma); err != nil {
			return err
		}

		sumas = append(sumas, suma)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSumaResponse{Suma: sumas, Pagination: pageRes}, nil
}

func (k Keeper) Suma(ctx context.Context, req *types.QueryGetSumaRequest) (*types.QueryGetSumaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	suma, found := k.GetSuma(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSumaResponse{Suma: suma}, nil
}
