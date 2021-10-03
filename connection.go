package solana

import "context"

// Connection represents a connection to a fullnode JSON RPC endpoint
type Connection interface {
	// Commitment returns the default commitmentLevel used for requests
	Commitment() CommitmentLevel

	// GetAccountInfo returns all the account info for the specified PublicKey
	GetAccountInfo(ctx context.Context, request GetAccountInfoRequest) (*GetAccountInfoResponse, error)

	// GetBalance returns the balance of the account of provided PublicKey
	GetBalance(ctx context.Context, request GetBalanceRequest) (*GetBalanceResponse, error)
}

type GetAccountInfoRequest struct {
	PublicKey        PublicKey `validate:"required"`
	CommitmentConfig CommitmentConfig
	Encoding         EncodingConfig
}

type GetAccountInfoResponse struct {
	Context Context
}

type GetBalanceRequest struct {
	PublicKey        PublicKey `validate:"required"`
	CommitmentConfig CommitmentConfig
}

type GetBalanceResponse struct {
	Context Context
	Value   uint64
}
