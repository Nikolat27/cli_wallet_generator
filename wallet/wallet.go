package wallet

import (
	"cli_wallet_generator/bip39"
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
