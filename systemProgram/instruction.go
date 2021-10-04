package systemProgram

// Instruction is a Solana system program Instruction.
// See rust defs here: https://github.com/solana-labs/solana/blob/4b2fe9b20d4c895f4d3cb58c2918c72a5b0a5b64/sdk/program/src/system_instruction.rs#L142
type Instruction uint32

const (
	CreateAccountInstruction Instruction = iota
	AssignInstruction
	TransferInstruction
	CreateAccountWithSeedInstruction
	AdvanceNonceAccountInstruction
	WithdrawNonceAccountInstruction
	InitializeNonceAccountInstruction
	AuthorizeNonceAccountInstruction
	AllocateInstruction
	AllocateWithSeedInstruction
	AssignWithSeedInstruction
	TransferWithSeedInstruction
)
