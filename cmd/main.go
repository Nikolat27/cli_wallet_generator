package main

import (
	"bufio"
	"cli_wallet_generator/cli"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter your commands: ")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			panic("Scan() is false")
		}

		input := strings.TrimSpace(scanner.Text())

		inputs := strings.Split(input, " ")
		if err := cli.HandleUserCommand(inputs); err != nil {
			fmt.Println(err)
		}
	}
}
