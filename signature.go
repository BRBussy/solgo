package solana

import (
	"github.com/BRBussy/solgo/internal/pkg/encoding"
)

// Signature is a digital signature in the ed25519 binary format and consumes 64 bytes
type Signature [64]byte

// Signatures is a list of Signature entries.
// It implements the encoding.Compactor interface so that
// it can be converted into an encoding.CompactArray of signatures.
type Signatures []Signature

// Compact Signatures into an encoding.CompactArray
func (s Signatures) Compact() (encoding.CompactArray, error) {
	return encoding.CompactArray{
		Length: 0,
		Data:   nil,
	}, nil
}
