package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateOwner{}

func NewMsgUpdateOwner(creator string, denom string, newOwner string) *MsgUpdateOwner {
	return &MsgUpdateOwner{
		Creator:  creator,
		Denom:    denom,
		NewOwner: newOwner,
	}
}

func (msg *MsgUpdateOwner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
