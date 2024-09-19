package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMintAndSendTokens{}

func NewMsgMintAndSendTokens(creator string, denom string, amount int32, recipient string) *MsgMintAndSendTokens {
	return &MsgMintAndSendTokens{
		Creator:   creator,
		Denom:     denom,
		Amount:    amount,
		Recipient: recipient,
	}
}

func (msg *MsgMintAndSendTokens) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
