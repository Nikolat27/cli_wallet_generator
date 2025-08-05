package main

import (
	"bufio"
	"fmt"
	"go_wallet_generator/cli"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if it exists (optional)
	godotenv.Load()

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
