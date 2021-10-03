package solana

import (
	"context"
	"github.com/BRBussy/solgo/internal/pkg/jsonrpc"
)

type jsonrpcConnectionConfig struct {
	network         Network
	endpoint        string
	commitmentLevel CommitmentLevel
}

type JSONRPCConnectionOption interface {
	apply(*jsonrpcConnectionConfig)
}

type jsonrpcConnectionOptionFunc func(*jsonrpcConnectionConfig)

func (fn jsonrpcConnectionOptionFunc) apply(cfg *jsonrpcConnectionConfig) {
	fn(cfg)
}

// WithCommitmentLevel sets CommitmentLevel on the JSONRPCConnection
func WithCommitmentLevel(c CommitmentLevel) JSONRPCConnectionOption {
	return jsonrpcConnectionOptionFunc(func(config *jsonrpcConnectionConfig) {
		config.commitmentLevel = c
	})
}

// WithNetwork sets Network on the JSONRPCConnection
// Note that this does not change the endpoint that the connection
// communicates with and so the WithEndpoint option may also need to
// be applied.
func WithNetwork(c Network) JSONRPCConnectionOption {
	return jsonrpcConnectionOptionFunc(func(config *jsonrpcConnectionConfig) {
		config.network = c
	})
}

// WithEndpoint sets endpoint on the JSONRPCConnection
func WithEndpoint(e string) JSONRPCConnectionOption {
	return jsonrpcConnectionOptionFunc(func(config *jsonrpcConnectionConfig) {
		config.endpoint = e
	})
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
		network:         MainnetBeta,
		endpoint:        MainnetBeta.MustToRPCURL(),
		commitmentLevel: ConfirmedCommitmentLevel,
	}

	// apply any provided options
	for _, opt := range opts {
		opt.apply(config)
	}

	return &JSONRPCConnection{
		jsonRPCClient: jsonrpc.NewHTTPClient(config.endpoint),
		config:        config,
	}
}

func (j *JSONRPCConnection) Commitment() CommitmentLevel {
	return j.config.commitmentLevel
}

func (j *JSONRPCConnection) GetBalance(ctx context.Context, request GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, nil
}
