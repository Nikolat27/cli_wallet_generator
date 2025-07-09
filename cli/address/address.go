package address

import (
	"cli_wallet_generator/address"
	"errors"
	"fmt"
)

const (
	walletNameFlag = "-w"
	coinNameFlag   = "-c"
)

func HandleAddressCommands(inputs []string) error {
	if len(inputs) < 2 {
		return errors.New("missing subcommand for address (e.g., create, get, delete, list) ⚠️ ")
	}

	switch cmd := inputs[1]; cmd{
	case "create":
		return handleCreate(inputs)
	default:
		return fmt.Errorf("unknown address subcommand: %s ❌  ", cmd)
	}
}

func handleCreate(inputs []string) error {
	if err := isCommandValid(inputs); err != nil {
		return err
	}

	walletName := inputs[3]
	coinName := inputs[5]

	_, err := address.GenerateAndStoreAddress(walletName, coinName)
	if err != nil {
		return err
	}

	return nil
}

func isCommandValid(inputs []string) error {
	if len(inputs) < 6 || inputs[2] != walletNameFlag || inputs[3] == "" || inputs[4] != coinNameFlag || inputs[5] == "" {
		return errors.New("invalid subcommand for coinAddress (e.g., create, get, delete, list) ⚠️ ")
	}

	return nil
}
