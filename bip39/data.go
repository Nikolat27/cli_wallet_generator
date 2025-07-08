package bip39

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	JsonFilePath = "wallets.json"
)

type WalletData struct {
	Wallets []Wallet `json:"wallets"`
}

type Wallet struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (s *Seed) storeWalletDataInJson(walletName string) (any, error) {
	data, err := os.ReadFile(JsonFilePath)
	if err != nil {
		return nil, err
	}

	// Handle empty file
	if len(data) == 0 {
		data = []byte(`{"wallets":[]}`)
	}

	var wallets WalletData
	if err := json.Unmarshal(data, &wallets); err != nil {
		return nil, err
	}
	
	if exist := wallets.checkWalletExist(walletName); exist {
		return "", fmt.Errorf("this wallet name exists already: %s", walletName)
	}

	wallets.Wallets = append(wallets.Wallets, Wallet{Name: walletName, Address: "testAddr"})

	updated, err := json.Marshal(wallets)
	if err != nil {
		return nil, err
	}

	if err := os.WriteFile(JsonFilePath, updated, 0644); err != nil {
		return nil, err
	}

	fmt.Println("Wallet added")
	return nil, nil
}

func (w *WalletData) checkWalletExist(name string) bool {
	for _, wallet := range w.Wallets {
		if wallet.Name == name {
			return true
		}
	}
	return false
}
