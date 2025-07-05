package seed

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

func (s *Seed) PrivateKeyToEthereumAddress() (string, error) {
	privateKey, err := crypto.ToECDSA(s.PrivateKey)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	pubBytes := crypto.FromECDSAPub(publicKey)
	
	pubBytes = pubBytes[1:]
	hash := crypto.Keccak256(pubBytes)

	address := hash[12:]

	finalAddress := "0x" + fmt.Sprintf("%x", address)
	return finalAddress, nil
}
