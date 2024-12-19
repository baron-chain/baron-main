// x/pqc/keeper/keeper.go
package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"baronchain/x/pqc/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
) *Keeper {
	// Ensure the module account is set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// StoreKyberPublicKey stores a public key for an address
func (k Keeper) StoreKyberPublicKey(ctx sdk.Context, address string, publicKey []byte) error {
	store := ctx.KVStore(k.storeKey)
	key := types.GetKyberPublicKeyKey(address)
	store.Set(key, publicKey)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeStoreKyberPublicKey,
			sdk.NewAttribute(types.AttributeKeyCreator, address),
			sdk.NewAttribute(types.AttributeKeyPublicKey, string(publicKey)),
		),
	)

	return nil
}

// GetKyberPublicKey retrieves a public key for an address
func (k Keeper) GetKyberPublicKey(ctx sdk.Context, address string) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetKyberPublicKeyKey(address)
	publicKey := store.Get(key)
	if publicKey == nil {
		return nil, fmt.Errorf("public key not found for address: %s", address)
	}
	return publicKey, nil
}

// SendEncryptedMessage stores an encrypted message for a recipient
func (k Keeper) SendEncryptedMessage(ctx sdk.Context, creator string, recipient string, ciphertext []byte) error {
	store := ctx.KVStore(k.storeKey)
	key := types.GetKyberCiphertextKey(recipient)
	store.Set(key, ciphertext)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSendEncryptedMessage,
			sdk.NewAttribute(types.AttributeKeyCreator, creator),
			sdk.NewAttribute(types.AttributeKeyRecipient, recipient),
			sdk.NewAttribute(types.AttributeKeyCiphertext, string(ciphertext)),
		),
	)

	return nil
}

// GetEncryptedMessage retrieves an encrypted message for a recipient
func (k Keeper) GetEncryptedMessage(ctx sdk.Context, recipient string) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetKyberCiphertextKey(recipient)
	ciphertext := store.Get(key)
	if ciphertext == nil {
		return nil, fmt.Errorf("no message found for recipient: %s", recipient)
	}
	return ciphertext, nil
}
