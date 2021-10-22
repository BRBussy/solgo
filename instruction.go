package solana

import (
	"github.com/BRBussy/solgo/internal/pkg/encoding"
)

// InstructionAccountMeta describes one of the accounts that will be provided
// as input to the program that is going to process an Instruction.
type InstructionAccountMeta struct {
	// PubKey is a public key identifying the account to be input
	PubKey PublicKey

	// IsSigner indicates if instruction requires the encompassing transaction
	// to contain signature for PubKey
	IsSigner bool

	// IsWritable indicates if the account should be loaded as a read-write account
	IsWritable bool
}

// Instruction is a Transaction instruction and is used to call some
// Program identified by ProgramIDPubKey
type Instruction struct {
	// InstructionAccountMeta is a slice of the accounts that are to
	// be provided to the entrypoint of the program being called.
	InstructionAccountMeta []InstructionAccountMeta

	// ProgramIDPubKey identifies the program to which this instruction
	// will be delivered.
	ProgramIDPubKey PublicKey

	// Data is the data to be input to the Program
	Data []byte
}

// Instructions is a list of Instruction entries.
// It implements the encoding.Compactor interface so that
// it can be converted into an encoding.CompactArray of signatures.
type Instructions []Instruction

func (i Instructions) Compact() (encoding.CompactArray, error) {
	panic("implement me")
}
