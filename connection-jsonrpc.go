package solana

import (
	"context"
	"github.com/BRBussy/solgo/internal/pkg/jsonrpc"
)

type jsonrpcConnectionConfig struct {
	network    Network
	endpoint   string
	commitment CommitmentLevel
}

type JSONRPCConnectionOption interface {
	apply(*jsonrpcConnectionConfig)
}

// JSONRPCConnection is a json-rpc http implementation of the solana.Connection interface
type JSONRPCConnection struct {
	jsonRPCClient jsonrpc.Client
	config        *jsonrpcConnectionConfig
}

// NewJSONRPCConnection returns a new and configured JSONRPCConnection.
//
// The default returned JSONRPCConnection is configured with:
//  - network: MainnetBeta
//  - commitmentLevel: ConfirmedCommitmentLevel
//
// The passed opts are used to override these default values and configure the
// returned JSONRPCConnection appropriately.
func NewJSONRPCConnection(opts ...JSONRPCConnectionOption) *JSONRPCConnection {
	// prepare default configuration
	config := &jsonrpcConnectionConfig{
		network:    MainnetBeta,
		endpoint:   MainnetBeta.MustToRPCURL(),
		commitment: ConfirmedCommitmentLevel,
	}

	return &JSONRPCConnection{
		jsonRPCClient: jsonrpc.NewHTTPClient(config.endpoint),
		config:        config,
	}
}

func (j *JSONRPCConnection) Commitment() CommitmentLevel {
	return j.config.commitment
}

func (j *JSONRPCConnection) GetBalance(ctx context.Context, request GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, nil
}
