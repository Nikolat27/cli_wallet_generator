package main

import (
	"fmt"
	"go_wallet_generator/frontend/app"
	"log"
	"net/http"
)

func main() {
	// Create and run the wallet application
	walletApp := app.NewWalletApp()

	port := ":3456"
	fmt.Printf("Starting wallet generator GUI at http://localhost%s\n", port)
	fmt.Println("Press Ctrl+C to stop the server")

	log.Fatal(http.ListenAndServe(port, walletApp))
}
