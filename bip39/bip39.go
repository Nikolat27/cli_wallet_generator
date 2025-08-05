package bip39

import (
	"github.com/tyler-smith/go-bip39"
)

type WalletSeed struct {
	Mnemonic []byte
	Seed     []byte
}

func InitWallet() (*WalletSeed, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return nil, err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}

	seed := bip39.NewSeed(mnemonic, "")

	w := &WalletSeed{
		Mnemonic: []byte(mnemonic),
		Seed:     seed,
	}

	if err := w.copyToClipboard(); err != nil {
		return nil, err
	}

	return w, nil
}

func (w *WalletSeed) Validate() bool {
	return bip39.IsMnemonicValid(string(w.Mnemonic))
}

func (w *WalletSeed) MnemonicString() string {
	return string(w.Mnemonic)
}

func (w *WalletSeed) SeedString() string {
	return string(w.Seed)
}

func (w *WalletSeed) copyToClipboard() error {
	// Don't display mnemonic here - it will be shown once in the CLI
	// This function is kept for compatibility but doesn't display anything
	return nil
}
