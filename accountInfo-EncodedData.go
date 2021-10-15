package solana

// AccountInfoEncodedData is information describing an account
// with data field encoded according to a prescribed Encoding
type AccountInfoEncodedData struct {
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

func (e AccountInfoEncodedData) GetExecutable() bool {
	return e.Executable
}

func (e AccountInfoEncodedData) GetLamports() uint64 {
	return e.Lamports
}

func (e AccountInfoEncodedData) GetOwner() string {
	return e.Owner
}

func (e AccountInfoEncodedData) GetRentEpoch() uint64 {
	return e.RentEpoch
}

func (e AccountInfoEncodedData) GetEncoding() Encoding {
	if len(e.Data) != 2 {
		return ""
	}
	return Encoding(e.Data[1])
}

func (e AccountInfoEncodedData) GetData() string {
	if len(e.Data) != 2 {
		return ""
	}
	return e.Data[0]
}
