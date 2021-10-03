package solana

// AccountInfo is information describing an account
type AccountInfo struct {
	// Executable is true if this account's Data contains a loaded program
	Executable bool `json:"executable"`

	// Lamports is the number of lamports assigned to this account
	Lamports uint64 `json:"lamports"`

	// Data is optional data assigned to the account
	Data []byte `json:"data"`

	// Owner is a base-58 encoded Pubkey of the program that owns this account
	Owner string `json:"owner"`

	// RentEpoch is the epoch at which this account will next owe rent
	RentEpoch uint64 `json:"rent_epoch"`
}

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

type EncodingConfig struct {
	Encoding AccountEncoding `json:"encoding"`
}

func (e EncodingConfig) IsBlank() bool {
	return e == EncodingConfig{}
}
