package solana

// FeeCalculator can be used to CalculateTransactionFee to send a given
// Transaction according to the fee schedule at this FeeScheduleBlockHash.
//
// Note: At the moment all this calculator's method CalculateTransactionFee does
// is just multiply the private property lamportsPerSignature by the no of signatures
// on a given transaction. The reason to abstract such a simple operation at this point
// into a FeeCalculator service provider of sorts is in attempt to model the Solana API
// GetRecentBlockHashResponse which has an object field named 'feeCalculator' that just
// contains the fee.
// With such a name it seems possible that this concept may be developed more at some point.
// i.e. perhaps at some point calculating the transaction fee may include more than just
// multiplying no of signatures by this number. At which point the object may have more fields
// added to it.
// Considering this - it may be a good idea to always call CalculateTransactionFee instead of
// using this FeeCalculator to get the LamportsPerSignature and multiplying by the no. of
// signatures on a transaction. This will avoid having to refactor if the calculation ever changes.
type FeeCalculator struct {
	blockHash string
	// lamportsPerSignature is essentially the 'fee schedule'.
	// This concept may be developed more at some point.
	// i.e. perhaps at some point calculating the transaction fee may
	// include more than just multiplying no of signatures by this number.
	// Note: If that ever happens it may prove useful to make FeeCalculator an
	// interface.
	lamportsPerSignature int64
}

// LamportsPerSignature is the amount of Lamports required per Transaction
// Signature according to the fee schedule for this FeeScheduleBlockHash.
func (f *FeeCalculator) LamportsPerSignature() int64 {
	return f.lamportsPerSignature
}

// FeeScheduleBlockHash is the hash of the block during which the
// fee schedule used by this FeeCalculator was retrieved.
func (f *FeeCalculator) FeeScheduleBlockHash() string {
	return f.blockHash
}

// CalculateTransactionFee determines the cost in Lamports to send the given Transaction.
func (f *FeeCalculator) CalculateTransactionFee(transaction Transaction) int64 {
	return f.lamportsPerSignature * int64(len(transaction.signatures))
}
