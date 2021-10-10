package encoding

// IntToCompactU16 encodes an integer into Compact-u16 format.
// A compact-u16 is a multi-byte encoding of 16 bits. The first byte contains the lower 7 bits
// of the value in its lower 7 bits. If the value is above 0x7f, the high bit is set and the next
// 7 bits of the value are placed into the lower 7 bits of a second byte. If the value is above
// 0x3fff, the high bit is set and the remaining 2 bits of the value are placed into the lower
// 2 bits of a third byte.
func IntToCompactU16() ([]byte, error) {
	return nil, nil
}
