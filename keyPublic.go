package solana

import (
	"crypto/ed25519"
	"github.com/btcsuite/btcutil/base58"
)

type PublicKey ed25519.PublicKey

func (p PublicKey) ToBase58() string {
	return base58.Encode(p)
}
