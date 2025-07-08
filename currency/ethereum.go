package address

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

type Ethereum struct {
	Address string
}

func (e *Ethereum) GetAddress(privKey []byte) error {
	privateKey, err := crypto.ToECDSA(privKey)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	pubBytes := crypto.FromECDSAPub(publicKey)

	pubBytes = pubBytes[1:]
	hash := crypto.Keccak256(pubBytes)

	Address := hash[12:]

	finalAddress := "0x" + fmt.Sprintf("%x", Address)

	e.Address = finalAddress
	return nil
}
