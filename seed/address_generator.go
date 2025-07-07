package seed

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

type Response struct {
	PrivKey string
	Address string
}

func (s *Seed) PrivateKeyToEthereumAddress() (*Response, error) {
	privateKey, err := crypto.ToECDSA(s.PrivateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	pubBytes := crypto.FromECDSAPub(publicKey)

	pubBytes = pubBytes[1:]
	hash := crypto.Keccak256(pubBytes)

	address := hash[12:]

	finalAddress := "0x" + fmt.Sprintf("%x", address)

	resp := &Response{
		PrivKey: fmt.Sprintf("%x", crypto.FromECDSA(privateKey)),
		Address: finalAddress,
	}

	return resp, nil
}
