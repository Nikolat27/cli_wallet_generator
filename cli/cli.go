package cli

import (
	"cli_wallet_generator/cli/address"
	"cli_wallet_generator/cli/wallet"
	"errors"
)

func HandleUserCommand(inputs []string) error {
	switch inputs[0] {
	case "wallet":
		return wallet.HandleWalletCommands(inputs)
	case "address":
		return address.HandleAddressCommands(inputs)
	default:
		return errors.New("invalid command")
	}
}
