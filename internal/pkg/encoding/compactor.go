package encoding

import "encoding/binary"

// Compactor is the interface implemented by types that
// can Compact themselves into a CompactArray
type Compactor interface {
	Compact() CompactArray
}

// CompactArray models the data structure described here:
// A compact-array is serialized as the array length, followed by each array item.
// Solana requires that the length be variant encoded over 3 bytes - i.e.:
// [ArrLenByte1, ArrLenByte2, ArrLenByte3, arrayContentsN, arrayContentsN+1,...]
// Source: https://docs.solana.com/developing/programming-model/transactions#compact-array-format
type CompactArray struct {
	Length uint64
	Data   []byte
}

// ToBytes gets the compact array as a byte array.
func (c CompactArray) ToBytes() []byte {
	// multi-byte variant encode the number of items in the array
	encodedArrayLength := make([]byte, binary.MaxVarintLen16)
	binary.PutUvarint(encodedArrayLength, c.Length)

	// and return the compact array as bytes
	return append(
		encodedArrayLength,
		c.Data...,
	)
}
