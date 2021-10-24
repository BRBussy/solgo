package solana

import (
	"github.com/BRBussy/solgo/internal/pkg/encoding"
)

// Signature is a digital signature in the ed25519 binary format and consumes 64 bytes
type Signature [64]byte

// Bytes returns the Signature as a byte slice instead of fixed size array.
func (s Signature) Bytes() []byte {
	return s[:]
}

// Signatures is a list of Signature entries.
// It implements the encoding.Compactor interface so that
// it can be converted into an encoding.CompactArray of signatures.
type Signatures []Signature

// Compact Signatures into an encoding.CompactArray
func (s Signatures) Compact() (encoding.CompactArray, error) {
	// prepare slice of data to return
	data := make([]byte, 0)

	// pack all signatures into the data slice
	for i := range s {
		data = append(data, s[i].Bytes()...)
	}

	// and return
	return encoding.CompactArray{
		Length: uint64(len(s)),
		Data:   data,
	}, nil
}
