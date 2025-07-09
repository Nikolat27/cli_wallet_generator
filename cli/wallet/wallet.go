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
	walletName, err := extractWalletName(inputs)
	if err != nil {
		return err
	}

	c.Wallet.Name = walletName

	if err := c.Wallet.CreateWallet(); err != nil {
		return err
	}

	fmt.Println("Wallet added successfully ✅  ")
	return nil
}

func (c *Commands) get(inputs []string) error {
	walletName, err := extractWalletName(inputs)
	if err != nil {
		return err
	}

	c.Wallet.Name = walletName

	w, err := c.Wallet.GetWalletInstance()
	if err != nil {
		return err
	}

	fmt.Printf("Wallet: %+v\n", *w)
	return nil
}

func (c *Commands) delete(inputs []string) error {
	walletName, err := extractWalletName(inputs)
	if err != nil {
		return err
	}

	c.Wallet.Name = walletName

	return c.Wallet.DeleteWallet()
}

func (c *Commands) list() error {
	wallets, err := c.Wallet.ListWallets()
	if err != nil {
		return err
	}

	for idx, w := range wallets {
		fmt.Printf("%d. Wallet: %s, created: %s\n", idx+1, w.Name, w.CreatedAt.Format("2006-01-02"))
	}
	return nil
}

func extractWalletName(inputs []string) (string, error) {
	if len(inputs) < 4 || inputs[2] != WalletNameFlag || inputs[3] == "" {
		return "", errors.New("missing or invalid wallet name flag")
	}
	return inputs[3], nil
}
