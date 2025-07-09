package currency

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	bip32 "github.com/tyler-smith/go-bip32"
)

type Ethereum struct {
	Mnemonic   string
}

const (
	Purpose      = 44
	CoinType     = 60
	Account      = 0
	Change       = 0
	AddressIndex = 0
)

// CreateAddress derives Ethereum address from mnemonic using BIP44 with go-bip32 lib
func (e *Ethereum) CreateAddress() (string, error) {
	// seed, err := bip39.NewSeed(e.Mnemonic)
	// if err != nil {
	// 	return "", err
	// }

	// masterKey, err := bip32.NewMasterKey(seed.RawSeed)
	// if err != nil {
	// 	return "", err
	// }

	// purposeKey, err := masterKey.NewChildKey(bip32.FirstHardenedChild + Purpose)
	// if err != nil {
	// 	return "", err
	// }

	// address, err := deriveBIP44(purposeKey)
	// if err != nil {
	// 	return "", err
	// }

	// return address, nil
	return "", nil
}

func deriveBIP44(purposeKey *bip32.Key) (string, error) {
	// Derive m/44'/60'
	coinTypeKey, err := purposeKey.NewChildKey(bip32.FirstHardenedChild + CoinType)
	if err != nil {
		return "", err
	}

	// Derive m/44'/60'/0'
	accountKey, err := coinTypeKey.NewChildKey(bip32.FirstHardenedChild + Account)
	if err != nil {
		return "", err
	}

	// Derive m/44'/60'/0'/0
	changeKey, err := accountKey.NewChildKey(Change)
	if err != nil {
		return "", err
	}

	// Derive m/44'/60'/0'/0/0
	addressKey, err := changeKey.NewChildKey(AddressIndex)
	if err != nil {
		return "", err
	}

	address, err := convertPrivateKeyToEthereumAddress(addressKey.Key)
	if err != nil {
		return "", err
	}

	return address, nil
}

func convertPrivateKeyToEthereumAddress(privateKeyBytes []byte) (string, error) {
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	pubBytes := crypto.FromECDSAPub(publicKey)[1:] // drop 0x04 prefix

	hash := crypto.Keccak256(pubBytes)
	address := hash[12:] // last 20 bytes

	finalAddress := "0x" + fmt.Sprintf("%x", address)
	return finalAddress, nil
}
