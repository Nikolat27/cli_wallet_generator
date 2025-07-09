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

func walletExists(wallets []Wallet, name string) bool {
	for _, wallet := range wallets {
		if wallet.Name == name {
			return true
		}
	}
	return false
}

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

// saveWallets -> Saves the wallets back to the jsonfile
func saveWallets(wallets []Wallet) error {
	data, err := json.MarshalIndent(wallets, "", "	")
	if err != nil {
		return err
	}
	return os.WriteFile(jsonFilePath, data, 0644)
}

// getWallets -> Public function to get wallets and checking if a name already exists
func getWallets(walletName string) ([]Wallet, error) {
	wallets, err := loadWallets()
	if err != nil {
		return nil, err
	}

	if walletExists(wallets, walletName) {
		return nil, fmt.Errorf("ERROR: wallet name '%s' already exists", walletName)
	}

	return wallets, nil
}

// AddWallet -> Adds a new wallet and saves it to the json file
func AddWallet(wallets []Wallet, name, encryptedMnemonic, address string) error {
	wallets = append(wallets, Wallet{
		Name:     name,
		Mnemonic: encryptedMnemonic,
		Address:  address,
	})

	if err := saveWallets(wallets); err != nil {
		return err
	}

	fmt.Println("Wallet Added Successfully")
	return nil
}

func getWallet(walletName string) (*Wallet, error) {
	wallets, err := loadWallets()
	if err != nil {
		return nil, err
	}

	if !walletExists(wallets, walletName) {
		return nil, fmt.Errorf("ERROR: wallet name '%s' does not exists", walletName)
	}

	for _, wallet := range wallets {
		if wallet.Name == walletName {
			return &wallet, nil
		}
	}

	return nil, nil
}

func deleteWallet(walletName string) error {
	wallets, err := loadWallets()
	if err != nil {
		return err
	}

	if !walletExists(wallets, walletName) {
		return fmt.Errorf("ERROR: wallet name '%s' does not exists", walletName)
	}

	for idx, wallet := range wallets {
		fmt.Println(idx, wallet)
	}
	
	return nil
}
