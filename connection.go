package solana

import "context"

// /Users/bernard/Projects/github.com/solana-labs/solana/client/src/rpc_client.rs

// Connection represents a connection to a fullnode JSON RPC endpoint
type Connection interface {
	// Commitment returns the default commitment used for requests
	Commitment(ctx context.Context) (Commitment, error)

	// GetBalance returns the balance of the account of provided PublicKey
	GetBalance(ctx context.Context, request GetBalanceRequest) (*GetBalanceResponse, error)
}

type GetBalanceRequest struct {
	PublicKey  PublicKey `validate:"required"`
	Commitment Commitment
}

type GetBalanceResponse struct {
	Context Context
	Balance uint64
}
