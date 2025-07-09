package cli

import (
	"cli_wallet_generator/cli/address"
	"cli_wallet_generator/cli/wallet"
	"errors"
)

func HandleUserCommand(inputs []string) error {
	firstCommand := inputs[0]
	switch firstCommand {
	case "wallet":
		return wallet.HandleWalletCommands(inputs)
	case "address":
		return address.HandleAddressCommands(inputs)
	default:
		return errors.New("invalid command")
	}
}
