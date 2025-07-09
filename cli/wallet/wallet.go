package wallet

import (
	"cli_wallet_generator/wallet"
	"errors"
	"fmt"
)

const (
	nameFlag = "--n"
)

func HandleWalletCommands(inputs []string) error {
	if len(inputs) < 2 {
		return errors.New("missing subcommand for wallet (e.g., create, get, delete, list) ⚠️ ")
	}

	command := inputs[1]
	switch command {
	case "create":
		return handleCreate(inputs)
	case "get":
		return handleGet(inputs)
	case "delete":
		return handleDelete(inputs)
	case "list":
		return handleList()
	default:
		return fmt.Errorf("unknown wallet subcommand: %s ❌  ", command)
	}
}

func handleCreate(inputs []string) error {
	if len(inputs) < 4 || inputs[2] != nameFlag {
		return errors.New("usage: wallet create --n <walletName> ⚠️ ")
	}
	walletName := inputs[3]
	return wallet.CreateWallet(walletName)
}

func handleGet(inputs []string) error {
	if len(inputs) < 4 || inputs[2] != nameFlag {
		return errors.New("usage: wallet get --n <walletName> ⚠️ ")
	}
	walletName := inputs[3]
	return wallet.GetWallet(walletName)
}

func handleDelete(inputs []string) error {
	if len(inputs) < 4 || inputs[2] != nameFlag {
		return errors.New("usage: wallet delete --n <walletName> ⚠️ ")
	}
	walletName := inputs[3]
	return wallet.DeleteWallet(walletName)
}

func handleList() error {
	return wallet.ListWallets()
}
