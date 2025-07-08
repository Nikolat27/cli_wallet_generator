package main

import (
	"cli_wallet_generator/bip39"
	address "cli_wallet_generator/currency"
)

func main() {
	mnemonic, err := bip39.NewMnemonic()
	if err != nil {
		panic(err)
	}

	seed, err := bip39.NewSeed(mnemonic.String())
	if err != nil {
		panic(err)
	}

	if err = seed.GenerateMasterKey(); err != nil {
		panic(err)
	}

	_ = address.InitCurrencies(seed.PrivateKey)
}
