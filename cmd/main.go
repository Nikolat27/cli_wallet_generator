package main

import (
	"cli_wallet_generator/bip39"
	"fmt"
)

func main() {
	seeds, err := bip39.GenerateSeeds()
	if err != nil {
		panic(err)
	}

	fmt.Println(seeds)
}
