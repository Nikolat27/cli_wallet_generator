package address

import (
	"cli_wallet_generator/crypto"
	"cli_wallet_generator/wallet"
	"errors"
	"fmt"
	"os"
	"time"
)

var coinsList = []string{"eth"}

func HandleCoinAddressGenerator(walletName, coinName string) (*wallet.Address, error) {
	walletInstance, err := wallet.GetWalletInstance(walletName)
	if err != nil {
		return nil, err
	}

	secretKey, err := getSecretKey()
	if err != nil {
		return nil, err
	}

	rawMnemonic, err := crypto.DecryptBase64([]byte(secretKey), walletInstance.Mnemonic)
	if err != nil {
		return nil, err
	}

	coinAddress, err := generateAddress(rawMnemonic, coinName)
	if err != nil {
		return nil, err
	}

	addressInstance := &wallet.Address{
		Coin:      coinName,
		Address:   coinAddress,
		CreatedAt: time.Now(),
	}

	walletInstance.AddAddress(addressInstance)

	wallets, err := wallet.LoadWallets()
	if err != nil {
		return nil, err
	}

	for i, w := range wallets {
		if w.Name == walletName {
			wallets[i] = *walletInstance
		}
	}

	err = wallet.SaveWallets(wallets)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func generateAddress(mnemonic []byte, coinName string) ([]byte, error) {
	switch coinName {
	case "eth":
		var eth = &Ethereum{
			RawMnemonic: mnemonic,
		}

		return eth.GenerateEthereumAddress(0)
	default:
		return nil, fmt.Errorf("ERROR invalid coin name: %s, available coins: %s", coinName, coinsList)
	}
}

func getSecretKey() (string, error) {
	key := os.Getenv("SECRET_KEY")
	if key == "" {
		return "", errors.New("SECRET_KEY env var is empty")
	}

	return key, nil
}
