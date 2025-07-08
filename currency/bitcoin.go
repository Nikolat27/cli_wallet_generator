package address

type Bitcoin struct {
	Address string
}

func (b *Bitcoin) GetAddress(privKey []byte) error {
	return nil
}
