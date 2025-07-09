package wallet

import (
	"cli_wallet_generator/bip39"
	"cli_wallet_generator/crypto"
	"fmt"
	"github.com/atotto/clipboard"
	"slices"
)

type Wallet struct {
	Name      string    `json:"name"`
	Mnemonic  string    `json:"mnemonic"`
	Addresses []Address `json:"addresses"`
}

// CreateWallet generates a new wallet, encrypts the mnemonic, and saves it
func CreateWallet(walletName string) error {
	wallets, err := LoadWallets()
	if err != nil {
		return fmt.Errorf("failed to load wallets: %w", err)
	}

	if IsWalletExist(wallets, walletName) {
		return fmt.Errorf("this wallet name '%s' already exists (must be unique)", walletName)
	}

	seed, err := bip39.InitWallet()
	if err != nil {
		return fmt.Errorf("failed to generate seed: %w", err)
	}

	if err := clipboard.WriteAll(seed.MnemonicString()); err != nil {
		return fmt.Errorf("failed to copy mnemonic to clipboard: %w", err)
	}
	fmt.Println("⚠️ 12-word seed copied to your clipboard. Keep it safe! ⚠️")

	encryptedMnemonic, err := crypto.EncryptBase64(seed.Mnemonic)
	if err != nil {
		return fmt.Errorf("failed to encrypt mnemonic: %w", err)
	}

	return addWallet(wallets, walletName, encryptedMnemonic, nil)
}

func ListWallets() error {
	wallets, err := LoadWallets()
	if err != nil {
		return fmt.Errorf("loading wallets failed: %w", err)
	}

	if len(wallets) == 0 {
		fmt.Println("No wallets found.")
		return nil
	}

	for idx, wallet := range wallets {
		fmt.Printf("Wallet %d: %s\n", idx+1, wallet.Name)
	}
	return nil
}

func GetWallet(walletName string) error {
	wallet, err := findWalletByName(walletName)
	if err != nil {
		return err
	}
	fmt.Printf("Wallet: %+v\n", *wallet)
	return nil
}

func GetWalletInstance(walletName string) (*Wallet, error) {
	wallet, err := findWalletByName(walletName)
	if err != nil {
		return nil, err
	}
	
	return wallet, nil
}

func DeleteWallet(walletName string) error {
	wallets, err := LoadWallets()
	if err != nil {
		return fmt.Errorf("failed to load wallets: %w", err)
	}

	idx := indexOfWallet(wallets, walletName)
	if idx == -1 {
		return fmt.Errorf("wallet '%s' does not exist", walletName)
	}

	wallets = slices.Delete(wallets, idx, idx+1)

	if err := SaveWallets(wallets); err != nil {
		return fmt.Errorf("failed to save wallets after deletion: %w", err)
	}

	fmt.Printf("Wallet '%s' deleted successfully ✅\n", walletName)
	return nil
}

func (w *Wallet) AddAddress(address *Address) {
	w.Addresses = append(w.Addresses, *address)
}
