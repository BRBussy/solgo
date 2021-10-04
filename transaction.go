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

// AddInstructions adds the given instructions to the transaction.
// An error will be returned if Signed is set.
func (t *Transaction) AddInstructions(i ...Instruction) (*Transaction, error) {
	// check if transaction has been signed
	if t.Signed {
		return nil, fmt.Errorf("cannot add instruction: %w", ErrTransactionAlreadySigned)
	}

	// instructions if not
	t.instructions = append(
		t.instructions,
		i...,
	)

	// and return the transaction
	return t, nil
}

// MustAddInstructions calls AddInstructions with the given instructions.
// Panics if AddInstructions returns an error - e.g. if Signed is set.
func (t *Transaction) MustAddInstructions(i ...Instruction) *Transaction {
	if _, err := t.AddInstructions(i...); err != nil {
		panic(err)
	}
	return t
}

func (t *Transaction) ToBase58() (string, error) {
	return "", nil
}

func (t *Transaction) ToBase64() (string, error) {
	return "", nil
}
