package wallet

import (
	"cli_wallet_generator/bip39"
	"errors"
	"fmt"

	"github.com/atotto/clipboard"
)

func CreateWallet(walletName string) error {
	wallets, err := getWallets(walletName)
	if err != nil {
		return fmt.Errorf("failed to load wallets: %w", err)
	}

	seed, err := bip39.Init()
	if err != nil {
		return fmt.Errorf("failed to generate seed: %w", err)
	}

	if err := clipboard.WriteAll(string(seed.Mnemonic)); err != nil {
		return fmt.Errorf("failed to copy mnemonic to clipboard: %w", err)
	}

	fmt.Println("⚠️  12-word seed copied to your clipboard. Keep it safe and private.")

	encryptedMnemonic, err := EncryptBase64(seed.Mnemonic)
	if err != nil {
		return fmt.Errorf("failed to save wallet: %w", err)
	}

	if err := AddWallet(wallets, walletName, encryptedMnemonic, ""); err != nil {
		return fmt.Errorf("failed to save wallet: %w", err)
	}

	return nil
}

func ListWallets() error {
	wallets, err := loadWallets()
	if err != nil {
		return fmt.Errorf("loading wallets error: %s", err)
	}

	for idx, wallet := range wallets {
		fmt.Printf("Wallet %d: %s \n", idx+1, wallet.Name)
	}

	return nil
}

func GetWallet(walletName string) error {
	wallet, err := getWallet(walletName)
	if err != nil {
		return err
	}

	if wallet == nil {
		return errors.New("no wallet exists with this name")
	}

	fmt.Println(wallet)

	return nil
}

func DeleteWallet(walletName string) error {
	if err := deleteWallet(walletName); err != nil {
		fmt.Errorf("ERROR deleting wallet: %s", err)
	}
	
	return nil
}
