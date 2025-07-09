package address

import (
	"cli_wallet_generator/crypto"
	"cli_wallet_generator/wallet"
	"errors"
	"fmt"
	"os"
	"time"
)

type DeriveAddressFunc func(wallet *wallet.Wallet) ([]byte, error)

var addressGenerators = map[string]DeriveAddressFunc{
	"btc": GenerateBitcoinAddress,
	"eth": GenerateEthereumAddress,
}

// GenerateAndStoreAddress -> Main func
func GenerateAndStoreAddress(walletName, coinName string) (*wallet.Address, error) {
	w, err := loadWalletWithMnemonic(walletName)
	if err != nil {
		return nil, err
	}

	addr, err := createAddressFromWallet(w, coinName)
	if err != nil {
		return nil, err
	}

	if err := updateWalletChanges(w); err != nil {
		return nil, err
	}

	return addr, nil
}

func loadWalletWithMnemonic(walletName string) (*wallet.Wallet, error) {
	w := wallet.Constructor()
	w.Name = walletName

	instance, err := w.GetWalletInstance()
	if err != nil {
		return nil, err
	}

	secretKey, err := getSecretKey()
	if err != nil {
		return nil, err
	}

	mnemonic, err := crypto.DecryptBase64([]byte(secretKey), instance.Mnemonic)
	if err != nil {
		return nil, err
	}

	instance.RawMnemonic = mnemonic
	return instance, nil
}

func createAddressFromWallet(w *wallet.Wallet, coinName string) (*wallet.Address, error) {
	coinAddress, err := generateCoinAddress(w, coinName)
	if err != nil {
		return nil, err
	}

	w.ClearRawMnemonic()

	addr := &wallet.Address{
		Coin:      coinName,
		Address:   string(coinAddress),
		CreatedAt: time.Now(),
	}

	w.AddAddress(addr)
	return addr, nil
}

func generateCoinAddress(w *wallet.Wallet, coin string) ([]byte, error) {
	generator, exists := addressGenerators[coin]
	if !exists {
		return nil, fmt.Errorf("unsupported coin: %s", coin)
	}
	return generator(w)
}

func updateWalletChanges(updated *wallet.Wallet) error {
	wallets, err := wallet.LoadFromDisk()
	if err != nil {
		return err
	}

	for i := range wallets {
		if wallets[i].Name == updated.Name {
			wallets[i] = *updated
			return wallet.SaveToDisk(wallets)
		}
	}

	return fmt.Errorf("wallet '%s' not found for update", updated.Name)
}

func getSecretKey() (string, error) {
	key := os.Getenv("SECRET_KEY")
	if key == "" {
		return "", errors.New("SECRET_KEY environment variable is missing")
	}
	return key, nil
}
