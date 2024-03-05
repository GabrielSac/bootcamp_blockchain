package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateSuma{}

func NewMsgCreateSuma(creator string, op1 int32, op2 int32, res int32) *MsgCreateSuma {
	return &MsgCreateSuma{
		Creator: creator,
		Op1:     op1,
		Op2:     op2,
		Res:	 res,
	}
}

func (msg *MsgCreateSuma) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSuma{}

func NewMsgUpdateSuma(creator string, id uint64, op1 int32, op2 int32, res int32) *MsgUpdateSuma {
	return &MsgUpdateSuma{
		Id:      id,
		Creator: creator,
		Op1:     op1,
		Op2:     op2,
		Res:     res,
	}
}

func (msg *MsgUpdateSuma) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteSuma{}

func NewMsgDeleteSuma(creator string, id uint64) *MsgDeleteSuma {
	return &MsgDeleteSuma{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteSuma) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
