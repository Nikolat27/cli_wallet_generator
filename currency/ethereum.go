package address

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

type Ethereum struct {
	PrivateKey []byte
}

func (e *Ethereum) GetAddress() (string, error) {
	privateKey, err := crypto.ToECDSA(e.PrivateKey)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	pubBytes := crypto.FromECDSAPub(publicKey)

	pubBytes = pubBytes[1:]
	hash := crypto.Keccak256(pubBytes)

	Address := hash[12:]

	finalAddress := "0x" + fmt.Sprintf("%x", Address)

	return finalAddress, nil
}
