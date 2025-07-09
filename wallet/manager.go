package wallet

import "fmt"

type Manager struct {
	wallets []Wallet
}

func NewManager() (*Manager, error) {
	wallets, err := LoadWallets()
	if err != nil {
		return nil, err
	}
	return &Manager{wallets: wallets}, nil
}

func (m *Manager) Save() error {
	return SaveWallets(m.wallets)
}

func (m *Manager) Find(name string) (*Wallet, error) {
	for i := range m.wallets {
		if m.wallets[i].Name == name {
			return &m.wallets[i], nil
		}
	}
	return nil, fmt.Errorf("wallet %s not found", name)
}

func (m *Manager) Update(w Wallet) error {
	for i := range m.wallets {
		if m.wallets[i].Name == w.Name {
			m.wallets[i] = w
			return m.Save()
		}
	}
	return fmt.Errorf("wallet %s not found", w.Name)
}

