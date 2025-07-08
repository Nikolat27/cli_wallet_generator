package currency

type Currency struct {
	Ethereum Ethereum
	Bitcoin  Bitcoin
}

func NewModels(mnemonic string) *Currency {
	return &Currency{
		Ethereum: Ethereum{mnemonic},
		Bitcoin:  Bitcoin{mnemonic},
	}
}
