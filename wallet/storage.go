package wallet

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const (
	JsonFilePath = "wallets.json"
	JsonIndent   = "    "
	JsonPrefix
)

// LoadWallets reads the wallet JSON file
func LoadWallets() ([]Wallet, error) {
	data, err := os.ReadFile(JsonFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Wallet{}, nil
		}
		return nil, err
	}

	if len(data) == 0 {
		return []Wallet{}, nil
	}

	var wallets []Wallet
	if err := json.Unmarshal(data, &wallets); err != nil {
		return nil, err
	}

	return wallets, nil
}

// SaveWallets writes the wallet list back to the json file
func SaveWallets(wallets []Wallet) error {
	data, err := json.MarshalIndent(wallets, JsonPrefix, JsonIndent)
	if err != nil {
		return err
	}
	return os.WriteFile(JsonFilePath, data, 0644)
}

func IsWalletExist(wallets []Wallet, name string) bool {
	return indexOfWallet(wallets, name) != -1
}

func indexOfWallet(wallets []Wallet, name string) int {
	for i, wallet := range wallets {
		if wallet.Name == name {
			return i
		}
	}
	return -1
}

func findWalletByName(name string) (*Wallet, error) {
	wallets, err := LoadWallets()
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

	if err := SaveWallets(wallets); err != nil {
		return err
	}

	fmt.Println("Wallet added successfully ✅  ")
	return nil
}
