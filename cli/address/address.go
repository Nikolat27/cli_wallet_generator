package address

import (
	"go_wallet_generator/address"
	"go_wallet_generator/wallet"
	"errors"
	"fmt"
)

const (
	walletNameFlag = "-w"
	coinNameFlag   = "-c"
)

type Commands struct {
	Address *wallet.Address
}

func HandleAddressCommands(inputs []string) error {
	if len(inputs) < 2 {
		return errors.New("missing subcommand for address (e.g., create, get, delete, list) ⚠️ ")
	}

	switch cmd := inputs[1]; cmd {
	case "create":
		return handleCreate(inputs)
	case "list":
		return handleList(inputs)
	default:
		return fmt.Errorf("unknown address subcommand: %s ❌  ", cmd)
	}
}

func handleCreate(inputs []string) error {
	if len(inputs) < 6 || inputs[2] != walletNameFlag || inputs[3] == "" || inputs[4] != coinNameFlag || inputs[5] == "" {
		return errors.New("invalid command. Usage: address create -w <walletName> -c <coinSymbol> ⚠️ ")
	}

	walletName := inputs[3]
	coinName := inputs[5]

	newAddress, err := address.GenerateAndStoreAddress(walletName, coinName)
	if err != nil {
		return err
	}

	fmt.Printf("%s address created successfully! ✅ \n", coinName)
	fmt.Printf("New Address: %s\n", newAddress.Address)
	
	return nil
}

func handleList(inputs []string) error {
	if len(inputs) < 4 || inputs[2] != walletNameFlag || inputs[3] == "" {
		return errors.New("invalid command. Usage: address list -w <walletName> ⚠️ ")
	}

	walletName := inputs[3]

	addresses, err := address.RetrieveAddressList(walletName)
	if err != nil {
		return err
	}

	for idx, addr := range addresses {
		fmt.Printf("%d. Coin: %s, Address: %s\n", idx+1, addr.Coin, addr.Address)
	}

	return nil
}
