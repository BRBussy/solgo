package encoding

// Compactor is the interface implemented by types that
// can Compact themselves into a CompactArray
type Compactor interface {
	Compact() (CompactArray, error)
}

// CompactArray models the data structure described here:
// https://docs.solana.com/developing/programming-model/transactions#compact-array-format
type CompactArray struct {
	Length uint64
	Data   []byte
}
