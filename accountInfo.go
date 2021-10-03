package solana

import "encoding/json"

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

// AccountInfoEncodedAccountData is information describing an account
// with data field encoded accorded to a prescribed AccountEncoding
type AccountInfoEncodedAccountData struct {
	// Executable is true if this account's Data contains a loaded program
	Executable bool `json:"executable"`

	// Lamports is the number of lamports assigned to this account
	Lamports uint64 `json:"lamports"`

	// Data is optional data assigned to the account
	Data []string `json:"data"`

	// Owner is a base-58 encoded Pubkey of the program that owns this account
	Owner string `json:"owner"`

	// RentEpoch is the epoch at which this account will next owe rent
	RentEpoch uint64 `json:"rentEpoch"`
}

func (e AccountInfoEncodedAccountData) GetEncoding() AccountEncoding {
	if len(e.Data) != 2 {
		return ""
	}
	return AccountEncoding(e.Data[1])
}

func (e AccountInfoEncodedAccountData) GetData() string {
	if len(e.Data) != 2 {
		return ""
	}
	return e.Data[0]
}

// AccountInfoJSONData is information describing an account
// with data field encoded accorded to a prescribed JSONParsedEncoding
type AccountInfoJSONData struct {
	// Executable is true if this account's Data contains a loaded program
	Executable bool `json:"executable"`

	// Lamports is the number of lamports assigned to this account
	Lamports uint64 `json:"lamports"`

	// Data is optional data assigned to the account
	Data map[string]map[string]json.RawMessage `json:"data"`

	// Owner is a base-58 encoded Pubkey of the program that owns this account
	Owner string `json:"owner"`

	// RentEpoch is the epoch at which this account will next owe rent
	RentEpoch uint64 `json:"rentEpoch"`
}
