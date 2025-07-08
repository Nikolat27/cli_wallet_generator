package address

type Currency struct {
	Ethereum Ethereum
	Bitcoin  Bitcoin
}

// func GenerateAddress(currency string) {
// 	switch currency {
// 	case "bitcoin":

// 	}
// }

func InitCurrencies(privKey []byte) *Currency {
	return &Currency{
		Ethereum: Ethereum{privKey},
		Bitcoin:  Bitcoin{privKey},
	}
}
