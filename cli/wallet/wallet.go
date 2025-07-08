package wallet

import (
	"cli_wallet_generator/bip39"
	"errors"
)

func HandleWalletCommands(inputs []string) error {
	if len(inputs) < 4 {
		return errors.New("command is short (expected another parameters)")
	}

	switch {
	// command => wallet create --name testName
 	case inputs[1] == "create" && inputs[2] == "--name" && inputs[3] != "":
		walletName := inputs[3]
		return bip39.CreateSeedAndMasterKey(walletName)
	default:
		return errors.New("invalid command")
	}
}
