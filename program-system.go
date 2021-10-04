package solana

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func init() {
	SystemProgram = &systemProgram{
		programID: NewPublicKeyFromBase58String("11111111111111111111111111111111"),
	}
}

// SystemProgram is set on module initialisation and can be used to
// construct SystemProgram instructions.
var SystemProgram *systemProgram

// SystemProgram is the api for the Solana system program.
// See instruction definitions here:
// https://github.com/solana-labs/solana/blob/4b2fe9b20d4c895f4d3cb58c2918c72a5b0a5b64/sdk/program/src/system_instruction.rs#L142
type systemProgram struct {
	programID PublicKey
}

type SystemProgramInstruction uint32

const (
	CreateAccountSystemProgramInstruction SystemProgramInstruction = iota
	AssignSystemProgramInstruction
	TransferSystemProgramInstruction
	CreateAccountWithSeedSystemProgramInstruction
	AdvanceNonceAccountSystemProgramInstruction
	WithdrawNonceAccountSystemProgramInstruction
	InitializeNonceAccountSystemProgramInstruction
	AuthorizeNonceAccountSystemProgramInstruction
	AllocateSystemProgramInstruction
	AllocateWithSeedSystemProgramInstruction
	AssignWithSeedSystemProgramInstruction
	TransferWithSeedSystemProgramInstruction
)

type CreateAccountParams struct {
	// FromPubkey is the account that will transfer the required Lamports
	// to cover the required Space to the new account
	// Req: [writer, signer]
	FromPubkey PublicKey

	// NewAccountPubkey is the public key for the new account
	// Req: [writer, signer]
	NewAccountPubkey PublicKey

	// Lamports is the amount of Lamports that will be transferred to the
	// new account on opening.
	Lamports uint64

	// Space is the amount of space in bytes to allocate to the new account
	Space uint64

	// ProgramID is the Public key of the program to assign as the owner of
	// the new account
	ProgramID PublicKey
}

type createAccountInstructionData struct {
	Instruction SystemProgramInstruction
	Lamports    uint64
	Space       uint64
	Owner       PublicKey
}

func (s *systemProgram) CreateAccount(params CreateAccountParams) (*Instruction, error) {
	// encode instruction data
	buf := new(bytes.Buffer)
	if err := binary.Write(
		buf,
		binary.LittleEndian,
		createAccountInstructionData{
			Instruction: CreateAccountSystemProgramInstruction,
			Lamports:    params.Lamports,
			Space:       params.Space,
			Owner:       params.ProgramID,
		},
	); err != nil {
		return nil, fmt.Errorf("error encoding create account data: %w", err)
	}

	// construct and return instruction
	return &Instruction{
		InstructionAccountMeta: []InstructionAccountMeta{
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
		ProgramIDPubKey: s.programID,
		Data:            buf.Bytes(),
	}, nil
}
