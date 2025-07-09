Install Dependencies:
    `go mod download`

Also install this (for clipboard):
    `sudo apt install xsel`

Run the project:
`make run` or `go run cmd/main.go`

Create Wallet:
wallet create -n <**walletName**>
`Example:
    wallet create -n w1`

Delete Wallet:
wallet delete -n <**walletName**>
`Example:
    wallet delete -n w1`

Get Wallet:
wallet get -n <**walletName**>
`Example:
    wallet get -n w1`

Wallet List:
`wallet list`

Create Address:
address create -w <**walletName**> -c <**coinSymbol**>
`Example:    
address create -w w1 -c btc`

Address List:
address list -w <**walletName**>
`Example:
address list -w w1`