package solana

import (
	"crypto/ed25519"
	"github.com/btcsuite/btcutil/base58"
)

type PrivateKey ed25519.PrivateKey

func NewPrivateKeyFromBase58String(privateKey string) PrivateKey {
	return base58.Decode(privateKey)
}

func (p PrivateKey) ToBase58() string {
	return base58.Encode(p)
}
