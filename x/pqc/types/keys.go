// x/pqc/types/keys.go
package types

const (
	// ModuleName defines the module name
	ModuleName = "pqc"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_pqc"
)

// KVStore keys
var (
	// KeyPrefixKyberPublicKey defines the prefix for storing Kyber public keys
	KeyPrefixKyberPublicKey = []byte{0x01}

	// KeyPrefixKyberCiphertext defines the prefix for storing Kyber ciphertexts
	KeyPrefixKyberCiphertext = []byte{0x02}
)

// GetKyberPublicKeyKey returns the store key to retrieve a KyberPublicKey from the index fields
func GetKyberPublicKeyKey(address string) []byte {
	return append(KeyPrefixKyberPublicKey, []byte(address)...)
}

// GetKyberCiphertextKey returns the store key to retrieve a KyberCiphertext from the index fields
func GetKyberCiphertextKey(address string) []byte {
	return append(KeyPrefixKyberCiphertext, []byte(address)...)
}
