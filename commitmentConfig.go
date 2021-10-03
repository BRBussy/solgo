package solana

type CommitmentConfig struct {
	CommitmentLevel CommitmentLevel
}

func (c CommitmentConfig) IsBlank() bool {
	return c == CommitmentConfig{}
}
