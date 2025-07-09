package wallet

import (
	"cli_wallet_generator/wallet"
	"errors"
	"fmt"
)

const (
	WalletNameFlag = "-n"
)

type Commands struct {
	Wallet *wallet.Wallet
}

func HandleWalletCommands(inputs []string) error {
	if len(inputs) < 2 {
		return errors.New("missing subcommand for wallet (e.g., create, get, delete, list) ⚠️ ")
	}

	c := &Commands{
		Wallet: wallet.Constructor(),
	}

	switch cmd := inputs[1]; cmd {
	case "create":
		return c.create(inputs)
	case "get":
		return c.get(inputs)
	case "delete":
		return c.delete(inputs)
	case "list":
		return c.list()
	default:
		return fmt.Errorf("unknown wallet subcommand: %s ❌ ", cmd)
	}
}

func (c *Commands) create(inputs []string) error {
	if len(inputs) < 4 || inputs[2] != WalletNameFlag {
		return errors.New("usage: wallet create --n <walletName> ⚠️ ")
	}

	c.Wallet.Name = inputs[3]

	if err := c.Wallet.CreateWallet(); err != nil {
		return err
	}

	fmt.Println("Wallet added successfully ✅  ")
	return nil
}

func (c *Commands) get(inputs []string) error {
	if len(inputs) < 4 || inputs[2] != WalletNameFlag {
		return fmt.Errorf("usage: wallet get %s <walletName>", WalletNameFlag)
	}

	c.Wallet.Name = inputs[3]

	w, err := c.Wallet.GetWalletInstance()
	if err != nil {
		return err
	}

	fmt.Printf("Wallet: %+v\n", *w)
	return nil
}

func (c *Commands) delete(inputs []string) error {
	if len(inputs) < 4 || inputs[2] != WalletNameFlag {
		return fmt.Errorf("usage: wallet delete %s <walletName> ⚠️ ", WalletNameFlag)
	}

	c.Wallet.Name = inputs[3]

	return c.Wallet.DeleteWallet()
}

func (c *Commands) list() error {
	wallets, err := c.Wallet.ListWallets()
	if err != nil {
		return err
	}

	for idx, w := range wallets {
		fmt.Printf("Wallet %d: %s\n", idx, w)
	}

	return nil
}
