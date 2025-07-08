package seed

import (
	"cli_wallet_generator/currency"
	"crypto/hmac"
	"crypto/pbkdf2"
	"crypto/sha512"
	"errors"
	"fmt"
)

type Seed struct {
	Bytes      []byte
	PrivateKey []byte
	ChainKey   []byte
	Currency   address.Currency
}

const (
	Pbkdf2Iterations = 2048
	Pbkdf2KeyLen     = 64
	PrivateKeyLength = 32
	ChainCodeLength  = 32
	bitcoinSeed      = "Bitcoin seed"
)

func NewSeed(mnemonic string) (*Seed, error) {
	salt, err := getSalt()
	if err != nil {
		return nil, err
	}

	s, err := pbkdf2.Key(sha512.New, mnemonic, []byte(salt), Pbkdf2Iterations, Pbkdf2KeyLen)
	if err != nil {
		return nil, err
	}

	seed := &Seed{
		Bytes: s,
	}

	return seed, nil
}

func getSalt() (string, error) {
	passphrase, err := getPassphrase()
	if err != nil {
		return "", nil
	}

	salt := "mnemonic" + passphrase

	return salt, nil
}

func getPassphrase() (string, error) {
	var userPassphrase string
	fmt.Print("Enter your passphrase (e.g.: apple, sam, trump and etc...): ")
	_, err := fmt.Scan(&userPassphrase)
	if err != nil {
		return "", err
	}

	return userPassphrase, nil
}

func (s *Seed) GenerateMasterKey() error {
	if len(s.Bytes) == 0 {
		return errors.New("seed bytes are empty")
	}

	hash := hmac.New(sha512.New, []byte(bitcoinSeed))
	if _, err := hash.Write(s.Bytes); err != nil {
		return fmt.Errorf("HMAC writing error: %s", err)
	}

	intermediateKey := hash.Sum(nil)
	if len(intermediateKey) != PrivateKeyLength+ChainCodeLength {
		return fmt.Errorf("invalid 'intermediateKey' length: %d", len(intermediateKey))
	}

	s.PrivateKey = intermediateKey[:32]
	s.ChainKey = intermediateKey[32:]

	return nil
}
