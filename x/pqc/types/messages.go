// x/pqc/types/messages.go
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgStoreKyberPublicKey  = "store_kyber_public_key"
	TypeMsgSendEncryptedMessage = "send_encrypted_message"
)

var _ sdk.Msg = &MsgStoreKyberPublicKey{}
var _ sdk.Msg = &MsgSendEncryptedMessage{}

// NewMsgStoreKyberPublicKey creates a new MsgStoreKyberPublicKey instance
func NewMsgStoreKyberPublicKey(creator string, publicKey []byte) *MsgStoreKyberPublicKey {
	return &MsgStoreKyberPublicKey{
		Creator:   creator,
		PublicKey: publicKey,
	}
}

func (msg *MsgStoreKyberPublicKey) Route() string {
	return RouterKey
}

func (msg *MsgStoreKyberPublicKey) Type() string {
	return TypeMsgStoreKyberPublicKey
}

func (msg *MsgStoreKyberPublicKey) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStoreKyberPublicKey) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.PublicKey) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "public key cannot be empty")
	}
	return nil
}

// NewMsgSendEncryptedMessage creates a new MsgSendEncryptedMessage instance
func NewMsgSendEncryptedMessage(creator string, recipient string, ciphertext []byte) *MsgSendEncryptedMessage {
	return &MsgSendEncryptedMessage{
		Creator:    creator,
		Recipient:  recipient,
		Ciphertext: ciphertext,
	}
}

func (msg *MsgSendEncryptedMessage) Route() string {
	return RouterKey
}

func (msg *MsgSendEncryptedMessage) Type() string {
	return TypeMsgSendEncryptedMessage
}

func (msg *MsgSendEncryptedMessage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendEncryptedMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address (%s)", err)
	}
	if len(msg.Ciphertext) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "ciphertext cannot be empty")
	}
	return nil
}

// Events
const (
	EventTypeStoreKyberPublicKey  = "store_kyber_public_key"
	EventTypeSendEncryptedMessage = "send_encrypted_message"

	AttributeKeyCreator    = "creator"
	AttributeKeyRecipient  = "recipient"
	AttributeKeyCiphertext = "ciphertext"
	AttributeKeyPublicKey  = "public_key"
)
