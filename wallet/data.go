package wallet

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const jsonFilePath = "wallets.json"

type Wallet struct {
	Name     string `json:"name"`
	Mnemonic string `json:"mnemonic"`
	Address  string `json:"address"`
}

// loadWallets reads the wallet JSON file
func loadWallets() ([]Wallet, error) {
	data, err := os.ReadFile(jsonFilePath)
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

// saveWallets writes the wallet list back to the json file
func saveWallets(wallets []Wallet) error {
	data, err := json.MarshalIndent(wallets, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(jsonFilePath, data, 0644)
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
	wallets, err := loadWallets()
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

func addWallet(wallets []Wallet, name, encryptedMnemonic, address string) error {
	wallets = append(wallets, Wallet{
		Name:     name,
		Mnemonic: encryptedMnemonic,
		Address:  address,
	})

	if err := saveWallets(wallets); err != nil {
		return err
	}

	fmt.Println("Wallet added successfully ✅  ")
	return nil
}
