package solana

// Transaction is a Solana blockchain transaction.
// Learn more at: https://docs.solana.com/developing/programming-model/transactions
type Transaction struct {
	// signatures is a list of digital signatures.
	// Each Digital Signature is in the ed25519 binary format and consumes 64 bytes.
	signatures Signatures // of Signatures

	// instructions is a list of Instructions.
	// Each Digital Signature is in the ed25519 binary format and consumes 64 bytes.
	instructions Instructions // of Instructions
}

// NewTransaction creates a new Transaction
func NewTransaction() *Transaction {
	return &Transaction{instructions: make(Instructions, 0)}
}

// AddInstructions adds the given instructions to the transaction.
// An error will be returned if the Transaction contains Signatures.
func (t *Transaction) AddInstructions(i ...Instruction) error {
	if len(t.signatures) > 0 {
		return ErrTransactionAlreadySigned
	}

	// instructions if not
	t.instructions = append(
		t.instructions,
		i...,
	)
	return nil
}

// Sign signs the Transaction with given PrivateKey(s) and appends
// a signature to a list of signatures held on the Transaction.
func (t *Transaction) Sign(pvtKeys ...PrivateKey) error {
	return nil
}

func (t *Transaction) ToBase58(recentBlockHash [32]byte) (string, error) {
	messageAccountAddresses := []byte{}

	// prepare compiled transaction
	compiledTxn := make([]byte, 0)
	for _, s := range [][]byte{
		// [1.] Compact array of signatures
		t.signatures.Compact().Data,

		// [2.] Message
		// [2.1] Header
		{
			// no. of required signatures in the contained transaction
			0x00,
			// no. of 'signed for' accounts that are read only
			0x00,
			// no. of read-only account addresses not requiring signatures
			0x00,
		},
		// [2.2] Account addresses
		messageAccountAddresses,
		// [2.3] Recent blockhash
		recentBlockHash[:],
		// [2.4] compact array of instructions
		t.instructions.Compact().Data,
	} {
		compiledTxn = append(compiledTxn, s...)
	}

	return "", nil
}

func (t *Transaction) ToBase64() (string, error) {
	return "", nil
}
