package solana

// Encoding specifies encoding for Data in Solana
type Encoding string

var (
	// Base58Encoding encodes the AccountInfo.Data field in base58 (is slower!)
	// With this encoding set account data must be < 129 bytes.
	Base58Encoding Encoding = "base58"

	// Base64Encoding encodes the AccountInfo.Data field in base64.
	// No restriction on account data size.
	Base64Encoding Encoding = "base64"

	// Base64PlusZSTDEncoding compresses the AccountInfo.Data using
	// Zstandard and base64 encodes the result.
	Base64PlusZSTDEncoding Encoding = "base64+zstd"

	// JSONParsedEncoding encoding attempts to use program-specific state
	// parsers to return more human-readable and explicit account state data.
	// If "jsonParsed" is requested but a parser cannot be found, the field
	// falls back to "base64" encoding, detectable when the data field is type <string>.
	JSONParsedEncoding Encoding = "jsonParsed"
)
