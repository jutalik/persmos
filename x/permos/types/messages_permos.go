package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePermos{}

func NewMsgCreatePermos(creator string, mypermos string) *MsgCreatePermos {
	return &MsgCreatePermos{
		Creator:  creator,
		Mypermos: mypermos,
	}
}

func (msg *MsgCreatePermos) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePermos{}

func NewMsgUpdatePermos(creator string, id uint64, mypermos string) *MsgUpdatePermos {
	return &MsgUpdatePermos{
		Id:       id,
		Creator:  creator,
		Mypermos: mypermos,
	}
}

func (msg *MsgUpdatePermos) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePermos{}

func NewMsgDeletePermos(creator string, id uint64) *MsgDeletePermos {
	return &MsgDeletePermos{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeletePermos) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
