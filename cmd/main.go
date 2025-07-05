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

	fmt.Println(string(newSeed.Bytes))
}
