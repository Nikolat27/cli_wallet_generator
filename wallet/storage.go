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

// LoadFromDisk -> reads the wallet JSON file
func LoadFromDisk() ([]Wallet, error) {
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

// SaveToDisk -> writes the wallet list back to the json file
func SaveToDisk(wallets []Wallet) error {
	data, err := json.MarshalIndent(wallets, JsonPrefix, JsonIndent)
	if err != nil {
		return fmt.Errorf("ERROR json marshaling the data: %s", err)
	}

	if err := os.WriteFile(JsonFilePath, data, 0644); err != nil {
		return fmt.Errorf("ERROR writing json file: %s", err)
	}

	return nil
}
