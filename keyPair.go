package solana

import (
	"crypto/ed25519"
	cryptoRand "crypto/rand"
)

type KeyPair struct {
	PublicKey
	PrivateKey
}

func NewRandomKeyPair() (*KeyPair, error) {
	pub, privateKey, err := ed25519.GenerateKey(cryptoRand.Reader)
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		PublicKey:  PublicKey(pub),
		PrivateKey: PrivateKey(privateKey),
	}, nil
}

func MustNewRandomKeypair() *KeyPair {
	pub, privateKey, err := ed25519.GenerateKey(cryptoRand.Reader)
	if err != nil {
		panic(err)
	}

	return &KeyPair{
		PublicKey:  PublicKey(pub),
		PrivateKey: PrivateKey(privateKey),
	}
}
