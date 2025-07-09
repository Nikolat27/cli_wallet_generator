Run the project:
make run

Create Wallet:
wallet create -n <**walletName**>
`Example:
    address create -n w1`

Delete Wallet:
wallet delete -n <**walletName**>
`Example:
    address delete -n w1`

Get Wallet:
wallet get -n <**walletName**>
`Example:
    address get -n w1`

Wallet List:
wallet list

Create Address:
address create -w <**walletName**> -c <**coinSymbol**>
`Example:    
address create -w w1 -c btc`

Address List:
address list -w <**walletName**>
`Example:
address list -w w1`