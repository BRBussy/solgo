package solana

import "fmt"

// Transaction is a Solana blockchain transaction.
// Learn more at: https://docs.solana.com/developing/programming-model/transactions
type Transaction struct {
	// Signed indicates that the transaction has been signed
	// at least once.
	Signed bool

	instructions []Instruction
}

// NewTransaction creates a new Transaction
func NewTransaction() *Transaction {
	return &Transaction{instructions: make([]Instruction, 0)}
}

// AddInstructions adds the given instructions to the transaction.
// An error will be returned if Signed is set.
func (t *Transaction) AddInstructions(i ...Instruction) error {
	// check if transaction has been signed
	if t.Signed {
		return fmt.Errorf("cannot add instructions to signed transaction: %w", ErrTransactionAlreadySigned)
	}

	// instructions if not
	t.instructions = append(
		t.instructions,
		i...,
	)

	return nil
}

func (t *Transaction) ToBase58() (string, error) {
	return "", nil
}

func (t *Transaction) ToBase64() (string, error) {
	return "", nil
}
