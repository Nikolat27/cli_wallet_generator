package wallet

import (
	"cli_wallet_generator/wallet"
	"errors"
	"fmt"
)

const (
	WalletNameFlag = "-n"
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
	if len(inputs) < 4 || inputs[2] != WalletNameFlag {
		return errors.New("usage: wallet create --n <walletName> ⚠️ ")
	}
	walletName := inputs[3]

	if err := wallet.CreateWallet(walletName); err != nil {
		return err
	}

	fmt.Println("Wallet added successfully ✅  ")
	return nil
}

func handleGet(inputs []string) error {
	if len(inputs) < 4 || inputs[2] != WalletNameFlag {
		return fmt.Errorf("usage: wallet get %s <walletName>", WalletNameFlag)
	}

	walletName := inputs[3]
	w, err := wallet.GetWalletInstance(walletName)
	if err != nil {
		return err
	}

	fmt.Printf("Wallet: %+v\n", *w)
	return nil
}

func handleDelete(inputs []string) error {
	if len(inputs) < 4 || inputs[2] != WalletNameFlag {
		return fmt.Errorf("usage: wallet delete %s <walletName> ⚠️ ", WalletNameFlag)
	}
	walletName := inputs[3]
	return wallet.DeleteWallet(walletName)
}

func handleList() error {
	wallets, err := wallet.ListWallets()
	if err != nil {
		return err
	}

	for idx, w := range wallets {
		fmt.Printf("Wallet %d: %s\n", idx, w)
	}

	return nil
}
