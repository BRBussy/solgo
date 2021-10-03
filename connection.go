package solana

import "context"

// /Users/bernard/Projects/github.com/solana-labs/solana/client/src/rpc_client.rs

// Connection represents a connection to a fullnode JSON RPC endpoint
type Connection interface {
	// Commitment returns the default commitmentLevel used for requests
	Commitment() CommitmentLevel

	// GetBalance returns the balance of the account of provided PublicKey
	GetBalance(ctx context.Context, request GetBalanceRequest) (*GetBalanceResponse, error)
}

type GetBalanceRequest struct {
	PublicKey  PublicKey `validate:"required"`
	Commitment CommitmentLevel
}

type GetBalanceResponse struct {
	Context Context
	Value   uint64
}
