package wallet

import (
	"cli_wallet_generator/bip39"
	"cli_wallet_generator/crypto"
	"fmt"
	"slices"
	"time"
)

type Wallet struct {
	Name      string    `json:"name"`
	Mnemonic  string    `json:"mnemonic"`
	Addresses []Address `json:"addresses"`
	RawMnemonic []byte
}

type Address struct {
	Coin      string    `json:"coin"`
	Address   []byte    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

func Constructor() *Wallet {
	return &Wallet{}
}

// CreateWallet generates a new wallet, encrypts the mnemonic, and saves it
func (w *Wallet) CreateWallet() error {
	wallets, err := LoadFromDisk()
	if err != nil {
		return fmt.Errorf("failed to load wallets: %w", err)
	}

	if exists := IsWalletExist(wallets, w.Name); exists {
		return fmt.Errorf("this wallet name '%s' already exists (must be unique)", w.Name)
	}

	seed, err := bip39.InitWallet()
	if err != nil {
		return fmt.Errorf("failed to generate seed: %w", err)
	}

	encryptedMnemonic, err := crypto.EncryptBase64(seed.Mnemonic)
	if err != nil {
		return fmt.Errorf("failed to encrypt mnemonic: %w", err)
	}

	if err := addWallet(wallets, w.Name, encryptedMnemonic, nil); err != nil {
		return err
	}

	return nil
}

func (w *Wallet) ListWallets() ([]Wallet, error) {
	wallets, err := LoadFromDisk()
	if err != nil {
		return nil, fmt.Errorf("loading wallets failed: %w", err)
	}

	return wallets, nil
}

func (w *Wallet) GetWalletInstance() (*Wallet, error) {
	wallet, err := findWalletByName(w.Name)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (w *Wallet) DeleteWallet() error {
	wallets, err := LoadFromDisk()
	if err != nil {
		return fmt.Errorf("failed to load wallets: %w", err)
	}

	idx := indexOfWallet(wallets, w.Name)
	if idx == -1 {
		return fmt.Errorf("wallet '%s' does not exist", w.Name)
	}

	wallets = slices.Delete(wallets, idx, idx+1)

	if err := SaveToDisk(wallets); err != nil {
		return fmt.Errorf("failed to save wallets after deletion: %w", err)
	}

	fmt.Printf("Wallet '%s' deleted successfully ✅\n", w.Name)
	return nil
}

func (w *Wallet) AddAddress(address *Address) {
	w.Addresses = append(w.Addresses, *address)
}

func (w *Wallet) ClearRawMnemonic() {
	w.RawMnemonic = nil
}

func IsWalletExist(wallets []Wallet, name string) bool {
	index := indexOfWallet(wallets, name)
	return index != -1
}

func indexOfWallet(wallets []Wallet, name string) int {
	for idx, wallet := range wallets {
		if wallet.Name == name {
			return idx
		}
	}
	return -1
}

func findWalletByName(name string) (*Wallet, error) {
	wallets, err := LoadFromDisk()
	if err != nil {
		return nil, err
	}

	for _, wallet := range wallets {
		if wallet.Name == name {
			return &wallet, nil
		}
	}
	return nil, fmt.Errorf("wallet '%s' not found", name)
}

func addWallet(wallets []Wallet, name, encryptedMnemonic string, addresses []Address) error {
	wallets = append(wallets, Wallet{
		Name:      name,
		Mnemonic:  encryptedMnemonic,
		Addresses: addresses,
	})

	if err := SaveToDisk(wallets); err != nil {
		return err
	}
	return nil
}
