package solana

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
