package address

import (
	"crypto/ecdsa"
	"fmt"
	"go_wallet_generator/wallet"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

const (
	Purpose          uint32 = 44
	EthereumCoinType uint32 = 60
	Change           uint32 = 0
	AddrIdx          uint32 = 0
)

type Ethereum struct {
	RawMnemonic []byte `json:"-"`
}

func GenerateEthereumAddress(wallet *wallet.Wallet) ([]byte, error) {
	eth := &Ethereum{RawMnemonic: wallet.RawMnemonic}

	addressIndex := eth.getNextAddressIndex(wallet)

	return eth.generateAddress(addressIndex)
}

func (e *Ethereum) generateAddress(addressIndex uint32) ([]byte, error) {
	seed := bip39.NewSeed(string(e.RawMnemonic), "")

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}

	purposeKey, err := masterKey.NewChildKey(bip32.FirstHardenedChild + Purpose)
	if err != nil {
		return nil, err
	}

	coinTypeKey, err := purposeKey.NewChildKey(bip32.FirstHardenedChild + EthereumCoinType)
	if err != nil {
		return nil, err
	}

	accountKey, err := coinTypeKey.NewChildKey(bip32.FirstHardenedChild + addressIndex)
	if err != nil {
		return nil, err
	}

	changeKey, err := accountKey.NewChildKey(Change)
	if err != nil {
		return nil, err
	}

	addressKey, err := changeKey.NewChildKey(AddrIdx)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.ToECDSA(addressKey.Key)
	if err != nil {
		return nil, err
	}

	address, err := deriveAddressFromPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	return address, nil
}

func deriveAddressFromPrivateKey(privateKey *ecdsa.PrivateKey) ([]byte, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	pubBytes := crypto.FromECDSAPub(publicKeyECDSA)[1:] // drop 0x04 prefix

	addressBytes := crypto.Keccak256(pubBytes)[12:] // last 20 bytes

	address := "0x" + fmt.Sprintf("%x", addressBytes)

	return []byte(address), nil
}

func (e *Ethereum) getName() string {
	return "eth"
}

func (e *Ethereum) getNextAddressIndex(wallet *wallet.Wallet) uint32 {
	var count uint32
	for _, addr := range wallet.Addresses {
		if addr.Coin == e.getName() {
			count++
		}
	}

	return count
}
