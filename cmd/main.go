package main

import (
	"cli_wallet_generator/bip39"
	"cli_wallet_generator/seed"
	"fmt"
)

func main() {
	mnemonic, err := bip39.NewMnemonic()
	if err != nil {
		panic(err)
	}

	newSeed, err := seed.NewSeed(mnemonic.String())
	if err != nil {
		panic(err)
	}

	if err = newSeed.GenerateMasterKey(); err != nil {
		panic(err)
	}

	if err = newSeed.Currency.Ethereum.GetAddress(newSeed.PrivateKey); err != nil {
		panic(err)
	}

	fmt.Println(newSeed.Currency.Bitcoin.Address)
}
