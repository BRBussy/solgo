package solana

func init() {
	SystemProgram = &systemProgram{
		programID: NewPublicKeyFromBase58String(""),
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

func (s *systemProgram) CreateAccount(params CreateAccountParams) (*Instruction, error) {
	return &Instruction{}, nil
}
