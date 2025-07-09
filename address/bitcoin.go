package address

import (
	"cli_wallet_generator/wallet"
	"crypto/sha256"
	"github.com/btcsuite/btcutil/base58"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/ripemd160"
)

const (
	BitcoinCoinType uint32 = 0
)

type Bitcoin struct {
	RawMnemonic []byte `json:"-"`
}

func GenerateBitcoinAddress(wallet *wallet.Wallet) ([]byte, error) {
	btc := &Bitcoin{RawMnemonic: wallet.RawMnemonic}

	addressIndex := btc.getNextAddressIndex(wallet)

	address, err := btc.generateAddress(addressIndex)
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (b *Bitcoin) generateAddress(addressIndex uint32) ([]byte, error) {
	seed := bip39.NewSeed(string(b.RawMnemonic), "")

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}

	// m / 44'
	purposeKey, err := masterKey.NewChildKey(bip32.FirstHardenedChild + Purpose)
	if err != nil {
		return nil, err
	}

	// m / 44' / 0'
	coinTypeKey, err := purposeKey.NewChildKey(bip32.FirstHardenedChild + BitcoinCoinType)
	if err != nil {
		return nil, err
	}

	// m / 44' / 0' / addressIndex'
	accountKey, err := coinTypeKey.NewChildKey(bip32.FirstHardenedChild + addressIndex)
	if err != nil {
		return nil, err
	}

	// m / 44' / 0' / addressIndex' / 0
	changeKey, err := accountKey.NewChildKey(Change)
	if err != nil {
		return nil, err
	}

	// m / 44' / 0' / addressIndex' / 0 / 0
	addressKey, err := changeKey.NewChildKey(AddrIdx)
	if err != nil {
		return nil, err
	}

	address, err := publicKeyToBitcoinAddress(addressKey.PublicKey().Key)
	if err != nil {
		return nil, err
	}

	return []byte(address), nil
}

func publicKeyToBitcoinAddress(pubKey []byte) (string, error) {
	shaHash := sha256.Sum256(pubKey)

	ripemd := ripemd160.New()
	if _, err := ripemd.Write(shaHash[:]); err != nil {
		return "", err
	}

	pubKeyHash := ripemd.Sum(nil)

	versionedPayload := append([]byte{0x00}, pubKeyHash...)

	firstSHA := sha256.Sum256(versionedPayload)
	secondSHA := sha256.Sum256(firstSHA[:])
	checksum := secondSHA[:4]

	finalPayload := append(versionedPayload, checksum...)

	address := base58.Encode(finalPayload)

	return address, nil
}

func (b *Bitcoin) getName() string {
	return "btc"
}

func (b *Bitcoin) getNextAddressIndex(wallet *wallet.Wallet) uint32 {
	var count uint32
	for _, addr := range wallet.Addresses {
		if addr.Coin == b.getName() {
			count++
		}
	}

	return count
}
