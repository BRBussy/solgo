package solana

import (
	"context"
	"fmt"
	"github.com/BRBussy/solgo/internal/pkg/jsonrpc"
	"github.com/BRBussy/solgo/internal/pkg/validation"
)

// JSONRPCConnection is a json-rpc http implementation of the solana.Connection interface
type JSONRPCConnection struct {
	validator     validation.Validator
	jsonRPCClient jsonrpc.Client
	config        *jsonrpcConnectionConfig
}

// jsonrpcConnectionConfig is the configuration for a JSONRPCConnection
type jsonrpcConnectionConfig struct {
	network         Network
	endpoint        string
	commitmentLevel CommitmentLevel
}

// JSONRPCConnectionOption makes a change to the jsonrpcConnectionConfig
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

// NewJSONRPCConnection returns a new and configured JSONRPCConnection.
//
// The default returned JSONRPCConnection is configured with:
//  - network: MainnetBeta
//  - endpoint: https://api.mainnet-beta.solana.com
//  - commitmentLevel: ConfirmedCommitmentLevel
//
// The passed opts are used to override these default values and configure the
// returned JSONRPCConnection as desired.
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

type GetBalanceJSONRPCResponse struct {
	Context Context `json:"context"`
	Value   uint64  `json:"value"`
}

func (j *JSONRPCConnection) GetBalance(ctx context.Context, request GetBalanceRequest) (*GetBalanceResponse, error) {
	// perform rpc call
	resp, err := j.jsonRPCClient.CallParamArray(
		ctx,
		"getBalance",
		nil,
		request.PublicKey.ToBase58(),
	)
	if err != nil {
		return nil, fmt.Errorf("error with getBalance json-rpc call: %w", err)
	}

	// parse response
	response := new(GetBalanceJSONRPCResponse)
	if err := resp.GetObject(response); err != nil {
		return nil, fmt.Errorf("error parsing GetBalanceJSONRPCResponse: %w", err)
	}

	return &GetBalanceResponse{
		Context: response.Context,
		Value:   response.Value,
	}, nil
}
