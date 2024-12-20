package kyber

import (
	"crypto/rand"
	"errors"

	kyberlib "github.com/cloudflare/circl/kem/kyber/kyber1024"
)

var (
	// Custom error types for detailed error handling
	ErrInvalidPublicKeySize  = errors.New("invalid public key size")
	ErrInvalidPrivateKeySize = errors.New("invalid private key size")
	ErrInvalidCiphertextSize = errors.New("invalid ciphertext size")
	ErrEncryptionFailed      = errors.New("key encapsulation failed")
	ErrDecryptionFailed      = errors.New("key decapsulation failed")
	ErrKeyGenerationFailed   = errors.New("key pair generation failed")
)

// Constants for Kyber-1024 parameters (NIST Level 5)
const (
	PublicKeySize  = kyberlib.PublicKeySize
	PrivateKeySize = kyberlib.PrivateKeySize
	CiphertextSize = kyberlib.CiphertextSize
	SharedKeySize  = kyberlib.SharedKeySize
)

// KeyPair represents a Kyber key pair
type KeyPair struct {
	publicKey  *kyberlib.PublicKey
	privateKey *kyberlib.PrivateKey
}

// NewKeyPair generates a new Kyber-1024 key pair
func NewKeyPair() (*KeyPair, error) {
	// Use system's cryptographically secure random source
	publicKey, privateKey, err := kyberlib.GenerateKeyPair(rand.Reader)
	if err != nil {
		return nil, ErrKeyGenerationFailed
	}

	return &KeyPair{
		publicKey:  publicKey,
		privateKey: privateKey,
	}, nil
}

// NewKeyFromSeed creates a key pair from a given seed
func NewKeyFromSeed(seed []byte) *KeyPair {
	publicKey, privateKey := kyberlib.NewKeyFromSeed(seed)
	return &KeyPair{
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

// PublicKeyBytes returns the public key as bytes
func (kp *KeyPair) PublicKeyBytes() []byte {
	publicKeyBytes := make([]byte, PublicKeySize)
	kp.publicKey.Pack(publicKeyBytes)
	return publicKeyBytes
}

// PrivateKeyBytes returns the private key as bytes
func (kp *KeyPair) PrivateKeyBytes() []byte {
	privateKeyBytes := make([]byte, PrivateKeySize)
	kp.privateKey.Pack(privateKeyBytes)
	return privateKeyBytes
}

// Encapsulate performs the key encapsulation mechanism
// Returns the shared secret and the ciphertext
func Encapsulate(publicKeyBytes []byte) (sharedSecret, ciphertext []byte, err error) {
	// Validate public key
	if len(publicKeyBytes) != PublicKeySize {
		return nil, nil, ErrInvalidPublicKeySize
	}

	// Unpack the public key
	publicKey := new(kyberlib.PublicKey)
	publicKey.Unpack(publicKeyBytes)

	// Prepare buffers for ciphertext and shared secret
	ct := make([]byte, CiphertextSize)
	ss := make([]byte, SharedKeySize)

	// Use nil seed for standard encapsulation
	publicKey.EncapsulateTo(ct, ss, nil)

	return ss, ct, nil
}

// Decapsulate performs the key decapsulation mechanism
// Returns the shared secret
func Decapsulate(privateKeyBytes, ciphertext []byte) ([]byte, error) {
	// Validate inputs
	if len(privateKeyBytes) != PrivateKeySize {
		return nil, ErrInvalidPrivateKeySize
	}
	if len(ciphertext) != CiphertextSize {
		return nil, ErrInvalidCiphertextSize
	}

	// Unpack the private key
	privateKey := new(kyberlib.PrivateKey)
	privateKey.Unpack(privateKeyBytes)

	// Prepare buffer for shared secret
	ss := make([]byte, SharedKeySize)

	// Decapsulate
	privateKey.DecapsulateTo(ss, ciphertext)

	return ss, nil
}

// SecureEncrypt provides a high-level encryption method
// Combines key encapsulation with optional additional encryption
func SecureEncrypt(publicKey, message []byte) (encapsulatedKey, encryptedMessage []byte, err error) {
	// Perform key encapsulation
	sharedSecret, ciphertext, err := Encapsulate(publicKey)
	if err != nil {
		return nil, nil, err
	}

	// XOR the message with the shared secret for additional security
	encryptedMsg := make([]byte, len(message))
	for i := range message {
		encryptedMsg[i] = message[i] ^ sharedSecret[i%SharedKeySize]
	}

	return ciphertext, encryptedMsg, nil
}

// SecureDecrypt provides a high-level decryption method
// Reverses the SecureEncrypt process
func SecureDecrypt(privateKey, ciphertext, encryptedMessage []byte) ([]byte, error) {
	// Decapsulate the shared secret
	sharedSecret, err := Decapsulate(privateKey, ciphertext)
	if err != nil {
		return nil, err
	}

	// Decrypt the message by XORing with the shared secret
	decryptedMsg := make([]byte, len(encryptedMessage))
	for i := range encryptedMessage {
		decryptedMsg[i] = encryptedMessage[i] ^ sharedSecret[i%SharedKeySize]
	}

	return decryptedMsg, nil
}
