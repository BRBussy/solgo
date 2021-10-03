package solana

import "context"

// /Users/bernard/Projects/github.com/solana-labs/solana/client/src/rpc_client.rs

// Connection represents a connection to a fullnode JSON RPC endpoint
type Connection interface {
	// Commitment returns the default commitment used for requests
	Commitment(ctx context.Context) (Commitment, error)

	GetBalance(
		ctx context.Context,
		publicKey PublicKey,
	) (*GetBalanceAndContextResponse, error)

	GetBalanceWithCommitment(
		ctx context.Context,
		publicKey PublicKey,
		commitment Commitment,
	) (*GetBalanceAndContextResponse, error)
}

type GetBalanceAndContextResponse struct {
	Context Context
	Balance uint64
}
