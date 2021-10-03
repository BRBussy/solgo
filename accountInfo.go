package solana

// AccountEncoding specifies encoding for AccountInfo.Data
type AccountEncoding string

var (
	// Base58AccountEncoding encodes the AccountInfo.Data field in base58 (is slower!)
	// With this encoding set account data must be < 129 bytes.
	Base58AccountEncoding AccountEncoding = "base58"

	// Base64AccountEncoding encodes the AccountInfo.Data field in base64.
	// No restriction on account data size.
	Base64AccountEncoding AccountEncoding = "base64"

	// Base64PlusZSTDAccountEncoding compresses the AccountInfo.Data using
	// Zstandard and base64 encodes the result.
	Base64PlusZSTDAccountEncoding AccountEncoding = "base64+zstd"

	// JSONParsedEncoding encoding attempts to use program-specific state
	// parsers to return more human-readable and explicit account state data.
	// If "jsonParsed" is requested but a parser cannot be found, the field
	// falls back to "base64" encoding, detectable when the data field is type <string>.
	JSONParsedEncoding AccountEncoding = "jsonParsed"
)

// AccountInfo is information describing an account
type AccountInfo interface {
	// GetExecutable returns true if this account's Data contains a loaded program
	GetExecutable() bool
	// GetLamports returns the number of lamports assigned to this account
	GetLamports() uint64
	// GetOwner returns a base58 encoded Pubkey of the program that owns this account
	GetOwner() string
	// GetRentEpoch returns the epoch at which this account will next owe rent
	GetRentEpoch() uint64
}
