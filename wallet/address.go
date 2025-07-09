package wallet

import "time"

type Address struct {
	Coin      string    `json:"coin"`
	Address   []byte    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}
