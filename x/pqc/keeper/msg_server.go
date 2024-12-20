// x/pqc/keeper/msg_server.go
package keeper

import (
	"context"
	"fmt"

	"baronchain/x/pqc/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// StoreKyberPublicKey handles the storage of a Kyber public key
func (k msgServer) StoreKyberPublicKey(goCtx context.Context, msg *types.MsgStoreKyberPublicKey) (*types.MsgStoreKyberPublicKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.Keeper.StoreKyberPublicKey(ctx, msg.Creator, msg.PublicKey)
	if err != nil {
		return nil, err
	}

	return &types.MsgStoreKyberPublicKeyResponse{}, nil
}

// SendEncryptedMessage handles the sending of an encrypted message
func (k msgServer) SendEncryptedMessage(goCtx context.Context, msg *types.MsgSendEncryptedMessage) (*types.MsgSendEncryptedMessageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.Keeper.SendEncryptedMessage(ctx, msg.Creator, msg.Recipient, msg.Ciphertext)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendEncryptedMessageResponse{}, nil
}

// UpdateParams handles updating the module parameters
func (k msgServer) UpdateParams(goCtx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if k.GetAuthority() != req.Authority {
		return nil, fmt.Errorf("invalid authority; expected %s, got %s", k.GetAuthority(), req.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	k.SetParams(ctx, req.Params)

	return &types.MsgUpdateParamsResponse{}, nil
}
