package address

import (
	"cli_wallet_generator/address"
	"errors"
	"fmt"
)

const (
	walletNameFlag = "--w"
	coinNameFlag   = "--c"
)

func HandleAddressCommands(inputs []string) error {
	if len(inputs) < 2 {
		return errors.New("missing subcommand for address (e.g., create, get, delete, list) ⚠️ ")
	}

	command := inputs[1]
	switch command {
	case "create":
		return handleCreate(inputs)
	default:
		return fmt.Errorf("unknown address subcommand: %s ❌  ", command)
	}
}

func handleCreate(inputs []string) error {
	if len(inputs) < 6 || inputs[2] != walletNameFlag || inputs[3] == "" || inputs[4] != coinNameFlag || inputs[5] == "" {
		return errors.New("invalid subcommand for coinAddress (e.g., create, get, delete, list) ⚠️ ")
	}

	walletName := inputs[3]
	coinName := inputs[5]

	_, err := address.HandleCoinAddressGenerator(walletName, coinName)
	if err != nil {
		return err
	}

	return nil
}
