package solana

type CommitmentConfig struct {
	Commitment CommitmentLevel `json:"commitment"`
}

func (c CommitmentConfig) IsBlank() bool {
	return c == CommitmentConfig{}
}
