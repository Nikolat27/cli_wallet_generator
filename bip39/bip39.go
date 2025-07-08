package bip39

import (
	bip39 "github.com/tyler-smith/go-bip39"
)

type WalletSeed struct {
	Mnemonic []byte
	Seed     []byte
}

func Init() (*WalletSeed, error) {
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
		Seed:     []byte(seed),
	}

	return w, nil
}

func (w *WalletSeed) Validate() bool {
	return bip39.IsMnemonicValid(string(w.Mnemonic))
}

func (w *WalletSeed) GetMnemonic() string {
	return string(w.Mnemonic)
}

func (w *WalletSeed) GetSeed() string {
	return string(w.Seed)
}
