package systemProgram

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/BRBussy/solgo"
)

type CreateAccountParams struct {
	// FromPubkey is the account that will transfer the required Lamports
	// to cover the required Space to the new account
	// Req: [writer, signer]
	FromPubkey solana.PublicKey

	// NewAccountPubkey is the public key for the new account
	// Req: [writer, signer]
	NewAccountPubkey solana.PublicKey

	// Lamports is the amount of Lamports that will be transferred to the
	// new account on opening.
	Lamports uint64

	// Space is the amount of space in bytes to allocate to the new account
	Space uint64

	// ProgramID is the Public key of the program to assign as the owner of
	// the new account
	ProgramID solana.PublicKey
}

type createAccountInstructionData struct {
	Instruction Instruction
	Lamports    uint64
	Space       uint64
	Owner       solana.PublicKey
}

// CreateAccount creates a Solana system program Instruction
func CreateAccount(params CreateAccountParams) (*solana.Instruction, error) {
	// encode instruction data
	buf := new(bytes.Buffer)
	if err := binary.Write(
		buf,
		binary.LittleEndian,
		createAccountInstructionData{
			Instruction: CreateAccountInstruction,
			Lamports:    params.Lamports,
			Space:       params.Space,
			Owner:       params.ProgramID,
		},
	); err != nil {
		return nil, fmt.Errorf("error encoding create account data: %w", err)
	}

	// construct and return instruction
	return &solana.Instruction{
		InstructionAccountMeta: []solana.InstructionAccountMeta{
			// 1st
			// Addresses requiring signatures are 1st, and in the following order:
			//
			// those that require write access
			{PubKey: params.FromPubkey, IsSigner: true, IsWritable: true},
			{PubKey: params.NewAccountPubkey, IsSigner: true, IsWritable: true},
			// those that require read-only access

			// 2nd
			// Addresses not requiring signatures are 2nd, and in the following order:
			//
			// those that require write access
			// those that require read-only access
		},
		ProgramIDPubKey: ID,
		Data:            buf.Bytes(),
	}, nil
}
