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
	// Transaction is the transaction being sent.
	// Note: it should already be signed.
	Transaction Transaction

	// SkipPreflight can be set to true to skip the
	// preflight transaction checks.
	// Default value if not specified is false.
	SkipPreflight bool

	// PreflightCommitmentLevel is the CommitmentLevel
	// to use for preflight checks.
	// Default value if not specified is "finalized".
	PreflightCommitmentLevel string

	// Encoding is the Encoding used for the transaction data.
	// Either "base58" (slow, DEPRECATED), or "base64".
	// Default value if not specified is "base58".
	Encoding Encoding

	// MaxRetries is them maximum number of times for the RPC node
	// to retry sending the transaction to the leader.
	// If this parameter not provided, the RPC node will retry the
	// transaction until it is finalized or until the blockhash expires.
	MaxRetries uint
}

type SendTransactionResponse struct {
	// TransactionID is the First Transaction Signature embedded
	// in the transaction, as base58 encoded string - aka. transaction id
	TransactionID string
}
