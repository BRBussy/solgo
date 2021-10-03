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

	// SendTransaction submits a signed transaction to the cluster for processing.
	// This method does not alter the transaction in any way; it relays the
	// transaction created by clients to the node as-is.
	SendTransaction(ctx context.Context, request SendTransactionRequest) (*SendTransactionResponse, error)
}

type GetAccountInfoRequest struct {
	PublicKey       PublicKey
	CommitmentLevel CommitmentLevel
	Encoding        Encoding
}

type GetAccountInfoResponse struct {
	Context     Context
	AccountInfo AccountInfo
}

type GetBalanceRequest struct {
	PublicKey       PublicKey
	CommitmentLevel CommitmentLevel
}

type GetBalanceResponse struct {
	Context Context
	Value   uint64
}

type SendTransactionRequest struct {
	CommitmentLevel CommitmentLevel
	// SkipPreflight can be set to true to skip the
	// preflight transaction checks.
	//  Default value if not specified is false.
	SkipPreflight bool

	// PreflightCommitment is the Commitment level to use for
	// preflight. Default value if not specified is "finalized".
	PreflightCommitment bool
}

type SendTransactionResponse struct {
}
