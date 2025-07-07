package seed

import (
	"crypto/hmac"
	"crypto/pbkdf2"
	"crypto/sha512"
	"fmt"
)

type Seed struct {
	Bytes      []byte
	PrivateKey []byte
	ChainKey   []byte
}

const (
	pbkdf2Iterations = 2048
	pbkdf2KeyLen     = 64
)

func NewSeed(mnemonic string) (*Seed, error) {
	salt, err := getSalt()
	if err != nil {
		return nil, err
	}

	s, err := pbkdf2.Key(sha512.New, mnemonic, []byte(salt), pbkdf2Iterations, pbkdf2KeyLen)
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

func (s *Seed) DeriveMasterKey() {
	mac := hmac.New(sha512.New, []byte("Bitcoin seed"))
	mac.Write(s.Bytes)

	I := mac.Sum(nil)

	s.PrivateKey = I[:32]
	s.ChainKey = I[32:]
}
