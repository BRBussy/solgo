package solana

import "context"

// Connection represents a connection to a fullnode JSON RPC endpoint
type Connection interface {
	// Commitment returns the default commitment used for requests
	Commitment(ctx context.Context) (Commitment, error)

	GetBalanceAndContext(
		ctx context.Context,
		publicKey PublicKey,
	) (*GetBalanceAndContextResponse, error)

	GetBalanceAndContextWithCommitment(
		ctx context.Context,
		publicKey PublicKey,
		commitment Commitment,
	) (*GetBalanceAndContextResponse, error)
}

type GetBalanceAndContextResponse struct {
}
