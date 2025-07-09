package main

import (
	"bufio"
	"cli_wallet_generator/cli"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func main() {
	if err := loadEnv(); err != nil {
		panic(err)
	}

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

func loadEnv() error {
	return godotenv.Load()
}
