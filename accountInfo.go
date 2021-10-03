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
