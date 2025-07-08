package currency

type Bitcoin struct {
	Mnemonic string
}

func (b *Bitcoin) GetAddress() error {
	return nil
}
