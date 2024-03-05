package keeper

import (
	"context"
	"fmt"

	"lambchain/x/lambchain/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateSuma(goCtx context.Context, msg *types.MsgCreateSuma) (*types.MsgCreateSumaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Op1 + msg.Op2 != msg.Res {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("%d + %d != %d", msg.Op1, msg.Op2, msg.Res))
	}

	var suma = types.Suma{
		Creator: msg.Creator,
		Op1:     msg.Op1,
		Op2:     msg.Op2,
		Res:     msg.Op1+msg.Op2,
	}

	id := k.AppendSuma(
		ctx,
		suma,
	)

	return &types.MsgCreateSumaResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateSuma(goCtx context.Context, msg *types.MsgUpdateSuma) (*types.MsgUpdateSumaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var suma = types.Suma{
		Creator: msg.Creator,
		Id:      msg.Id,
		Op1:     msg.Op1,
		Op2:     msg.Op2,
		Res:     msg.Res,
	}

	// Checks that the element exists
	val, found := k.GetSuma(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetSuma(ctx, suma)

	return &types.MsgUpdateSumaResponse{}, nil
}

func (k msgServer) DeleteSuma(goCtx context.Context, msg *types.MsgDeleteSuma) (*types.MsgDeleteSumaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetSuma(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveSuma(ctx, msg.Id)

	return &types.MsgDeleteSumaResponse{}, nil
}
