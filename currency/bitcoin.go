package address

type Bitcoin struct {
	PrivateKey []byte
}

func (b *Bitcoin) GetAddress() error {
	return nil
}
