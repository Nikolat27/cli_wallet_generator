package wallet

import (
	"cli_wallet_generator/wallet"
	"errors"
)

func HandleWalletCommands(inputs []string) error {
	if len(inputs) < 2 {
		return errors.New("command is short (expected another parameters)")
	}

	switch {
	// command => wallet create --name testName
	case inputs[1] == "create" && inputs[2] == "--name" && inputs[3] != "":
		walletName := inputs[3]
		return wallet.CreateWallet(walletName)
	case inputs[1] == "list":
		return wallet.ListWallets()
	case inputs[1] == "get" && inputs[2] == "--name" && inputs[3] != "":
		walletName := inputs[3]
		return wallet.GetWallet(walletName)
	case inputs[1] == "delete" && inputs[2] == "--name" && inputs[3] != "":
		walletName := inputs[3]
		return wallet.DeleteWallet(walletName)
	default:
		return errors.New("invalid command")
	}
}
