package solana

import "encoding/json"

// AccountInfoJSONData is information describing an account
// with data field encoded accorded to a prescribed JSONParsedEncoding
type AccountInfoJSONData struct {
	// Executable is true if this account's Data contains a loaded program
	Executable bool `json:"executable"`

	// Lamports is the number of lamports assigned to this account
	Lamports uint64 `json:"lamports"`

	// Data is optional data assigned to the account
	Data map[string]json.RawMessage `json:"data"`

	// Owner is a base58 encoded Pubkey of the program that owns this account
	Owner string `json:"owner"`

	// RentEpoch is the epoch at which this account will next owe rent
	RentEpoch uint64 `json:"rentEpoch"`
}

func (a AccountInfoJSONData) GetExecutable() bool {
	return a.Executable
}

func (a AccountInfoJSONData) GetLamports() uint64 {
	return a.Lamports
}

func (a AccountInfoJSONData) GetOwner() string {
	return a.Owner
}

func (a AccountInfoJSONData) GetRentEpoch() uint64 {
	return a.RentEpoch
}
