package solana

import (
	"crypto/ed25519"
	"github.com/btcsuite/btcutil/base58"
)

type PrivateKey struct {
	ed25519.PrivateKey
}

func NewPrivateKeyFromBase58String(privateKey string) PrivateKey {
	return PrivateKey{PrivateKey: base58.Decode(privateKey)}
}

func (p PrivateKey) ToBase58() string {
	return base58.Encode(p.PrivateKey)
}

func (p PrivateKey) PublicKey() PublicKey {
	ed25519PublicKey, ok := p.PrivateKey.Public().(ed25519.PublicKey)
	if !ok {
		panic("unable to infer ed25519 public type from private key")
	}
	return PublicKey{PublicKey: ed25519PublicKey}
}
