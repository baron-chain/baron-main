// x/pqc/keeper/query_server.go
package keeper

import (
	"context"

	"baronchain/x/pqc/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type queryServer struct {
	Keeper
}

// NewQueryServerImpl returns an implementation of the QueryServer interface
// for the provided Keeper.
func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return &queryServer{Keeper: keeper}
}

var _ types.QueryServer = queryServer{}

// KyberPublicKey implements the Query/KyberPublicKey gRPC method
func (k queryServer) KyberPublicKey(goCtx context.Context, req *types.QueryKyberPublicKeyRequest) (*types.QueryKyberPublicKeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	publicKey, err := k.Keeper.GetKyberPublicKey(ctx, req.Address)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &types.QueryKyberPublicKeyResponse{
		PublicKey: publicKey,
	}, nil
}

// EncryptedMessage implements the Query/EncryptedMessage gRPC method
func (k queryServer) EncryptedMessage(goCtx context.Context, req *types.QueryEncryptedMessageRequest) (*types.QueryEncryptedMessageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	ciphertext, err := k.Keeper.GetEncryptedMessage(ctx, req.Recipient)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &types.QueryEncryptedMessageResponse{
		Ciphertext: ciphertext,
	}, nil
}
