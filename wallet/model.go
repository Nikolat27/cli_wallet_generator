package wallet

type Wallet struct {
	Name     string `json:"name"`
	Mnemonic string `json:"mnemonic"`
	Address  string `json:"address"`
}
