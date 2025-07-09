package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/zalando/go-keyring"
	"io"
)

const (
	keyringService = "my_wallet_cli"
	keyringUser    = "wallet_secret_key"
	keySize        = 32 // 256 bits
)

func getOrCreateSecretKey() ([]byte, error) {
	storedKey, err := keyring.Get(keyringService, keyringUser)
	if err == nil {
		// Key exists, decode and return
		keyBytes, err := base64.StdEncoding.DecodeString(storedKey)
		if err != nil {
			return nil, fmt.Errorf("failed to decode stored key: %w", err)
		}
		
		return keyBytes, nil
	}

	if !errors.Is(err, keyring.ErrNotFound) {
		return nil, fmt.Errorf("keyring get error: %w", err)
	}

	// generate a new random key
	key := make([]byte, keySize)
	if _, err := rand.Read(key); err != nil {
		return nil, fmt.Errorf("failed to generate random key: %w", err)
	}

	// base64 encoded
	keyEncoded := base64.StdEncoding.EncodeToString(key)
	if err := keyring.Set(keyringService, keyringUser, keyEncoded); err != nil {
		return nil, fmt.Errorf("failed to set key in keyring: %w", err)
	}

	return key, nil
}

// EncryptBase64 -> Encryption and 64base encoding
func EncryptBase64(plaintext []byte) (string, error) {
	key, err := getOrCreateSecretKey()
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
	full := append(nonce, ciphertext...)

	return base64.StdEncoding.EncodeToString(full), nil
}

// DecryptBase64 -> Decryption and 64base decoding
func DecryptBase64(b64data string) ([]byte, error) {
	key, err := getOrCreateSecretKey()
	if err != nil {
		return nil, fmt.Errorf("failed to get secret key: %w", err)
	}

	rawData, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		return nil, fmt.Errorf("base64 decode error: %w", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher block: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM cipher: %w", err)
	}

	if len(rawData) < gcm.NonceSize() {
		return nil, errors.New("ciphertext too short")
	}

	nonce := rawData[:gcm.NonceSize()]
	ciphertext := rawData[gcm.NonceSize():]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("decryption error: %w", err)
	}

	return plaintext, nil
}
