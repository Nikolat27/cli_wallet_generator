package bip39

import (
	"crypto/hmac"
	"crypto/pbkdf2"
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/zalando/go-keyring"
)

type Seed struct {
	RawSeed    []byte
	PrivateKey []byte
	ChainKey   []byte
}

const (
	Pbkdf2Iterations = 2048
	Pbkdf2KeyLen     = 64
	PrivateKeyLength = 32
	ChainCodeLength  = 32
	bitcoinSeed      = "Bitcoin seed"
	ServiceName      = "cli-wallet-generator"
	KeyringUsername  = "832123"
	SecretKey        = "bmqo1k489sklz!r2"
)

func CreateSeedAndMasterKey(walletName string) error {
	mnemonic, err := newMnemonic()
	if err != nil {
		return err
	}

	seed, err := NewSeed(mnemonic.String())
	if err != nil {
		return err
	}

	if err := seed.generateMasterKey(); err != nil {
		return err
	}

	if _, err := seed.storeWalletDataInJson(walletName); err != nil {
		return err
	}

	fmt.Printf("\n12 phrase words Copied to your Clipboard (Don`t ever expose it)\n")
	return clipboard.WriteAll(mnemonic.String())
}

// NewSeed => Convert 12 words to 64-byte Seed using pbkdf2
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
		RawSeed: s,
	}

	return seed, nil
}

func getSalt() (string, error) {
	passphrase, err := getPassphrase()
	if err != nil {
		return "", err
	}

	salt := "mnemonic" + passphrase

	return salt, nil
}

func getPassphrase() (string, error) {
	var userPassphrase string
	fmt.Print("Enter your passphrase (e.g.: apple, sam, trump and etc...): ")

	if _, err := fmt.Scan(&userPassphrase); err != nil {
		return "", err
	}

	return userPassphrase, nil
}

func (s *Seed) generateMasterKey() error {
	if len(s.RawSeed) == 0 {
		return errors.New("seed RawSeed are empty")
	}

	hash := hmac.New(sha512.New, []byte(bitcoinSeed))
	if _, err := hash.Write(s.RawSeed); err != nil {
		return fmt.Errorf("HMAC writing error: %s", err)
	}

	I := hash.Sum(nil)
	if len(I) != PrivateKeyLength+ChainCodeLength {
		return fmt.Errorf("invalid 'I' length: %d", len(I))
	}

	s.PrivateKey = I[:PrivateKeyLength]
	s.ChainKey = I[PrivateKeyLength:]

	return nil
}

func (s *Seed) storeSeedInKeyring(user string) error {
	return keyring.Set(ServiceName, user, string(s.RawSeed))
}
