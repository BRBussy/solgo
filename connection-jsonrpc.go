package solana

import (
	"context"
	"fmt"
	"github.com/BRBussy/solgo/internal/pkg/jsonrpc"
)

// ensure JSONRPCConnection implements Connection
var _ Connection = &JSONRPCConnection{}

// JSONRPCConnection is a json-rpc http implementation of the solana.Connection interface
type JSONRPCConnection struct {
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

func (j *JSONRPCConnection) GetAccountInfo(ctx context.Context, request GetAccountInfoRequest) (*GetAccountInfoResponse, error) {
	// prepare configuration object
	config := map[string]interface{}{
		"commitment": j.Commitment(),
		"encoding":   Base64Encoding,
	}

	// set commitment level if provided
	if request.CommitmentLevel != "" {
		config["commitment"] = request.CommitmentLevel
	}

	// set encoding if provided
	if request.Encoding != "" {
		config["encoding"] = request.Encoding
	}

	// perform rpc call
	rpcResponse, err := j.jsonRPCClient.CallParamArray(
		ctx,
		"getAccountInfo",
		nil,
		request.PublicKey.ToBase58(),
		config,
	)
	if err != nil {
		return nil, fmt.Errorf("error performing getAccountInfo json-rpc call: %w", err)
	}
	if rpcResponse.Error != nil {
		return nil, fmt.Errorf("error set on rpc response: %w", rpcResponse.Error)
	}

	// parse response by type
	var response GetAccountInfoResponse
	switch request.Encoding {
	case JSONParsedEncoding:
		r := new(
			struct {
				Context Context             `json:"context"`
				Value   AccountInfoJSONData `json:"value"`
			},
		)
		if err := rpcResponse.GetObject(r); err != nil {
			return nil, fmt.Errorf("error parsing getBalanceJSONRPCResponse: %w", err)
		}
		response.Context = r.Context
		response.AccountInfo = r.Value

	default:
		r := new(
			struct {
				Context Context                `json:"context"`
				Value   AccountInfoEncodedData `json:"value"`
			},
		)
		if err := rpcResponse.GetObject(r); err != nil {
			return nil, fmt.Errorf("error parsing getBalanceJSONRPCResponse: %w", err)
		}
		response.Context = r.Context
		response.AccountInfo = r.Value
	}

	return &response, nil
}

type getBalanceJSONRPCResponse struct {
	Context Context `json:"context"`
	Value   uint64  `json:"value"`
}

func (j *JSONRPCConnection) GetBalance(ctx context.Context, request GetBalanceRequest) (*GetBalanceResponse, error) {
	// prepare configuration object
	config := map[string]interface{}{
		"commitment": j.Commitment(),
	}

	// set commitment level if provided
	if request.CommitmentLevel != "" {
		config["commitment"] = request.CommitmentLevel
	}

	// perform rpc call
	rpcResponse, err := j.jsonRPCClient.CallParamArray(
		ctx,
		"getBalance",
		nil,
		request.PublicKey.ToBase58(),
		config,
	)
	if err != nil {
		return nil, fmt.Errorf("error performing getBalance json-rpc call: %w", err)
	}
	if rpcResponse.Error != nil {
		return nil, fmt.Errorf("error set on rpc response: %w", rpcResponse.Error)
	}

	// parse response
	response := new(getBalanceJSONRPCResponse)
	if err := rpcResponse.GetObject(response); err != nil {
		return nil, fmt.Errorf("error parsing getBalanceJSONRPCResponse: %w", err)
	}

	return &GetBalanceResponse{
		Context: response.Context,
		Value:   response.Value,
	}, nil
}

type getRecentBlockHashJSONRPCResponse struct {
	Context Context `json:"context"`
	Value   struct {
		BlockHash     string `json:"blockhash"`
		FeeCalculator struct {
			LamportsPerSignature int64 `json:"lamportsPerSignature"`
		} `json:"feeCalculator"`
	}
}

func (j *JSONRPCConnection) GetRecentBlockHash(ctx context.Context, request GetRecentBlockHashRequest) (*GetRecentBlockHashResponse, error) {
	// prepare configuration object
	config := map[string]interface{}{
		"commitment": j.Commitment(),
	}

	// set commitment level if provided
	if request.CommitmentLevel != "" {
		config["commitment"] = request.CommitmentLevel
	}

	// perform rpc call
	rpcResponse, err := j.jsonRPCClient.CallParamArray(
		ctx,
		"getRecentBlockhash",
		nil,
		config,
	)
	if err != nil {
		return nil, fmt.Errorf("error performing getRecentBlockhash json-rpc call: %w", err)
	}
	if rpcResponse.Error != nil {
		return nil, fmt.Errorf("error set on rpc response: %w", rpcResponse.Error)
	}

	// parse response
	response := new(getRecentBlockHashJSONRPCResponse)
	if err := rpcResponse.GetObject(response); err != nil {
		return nil, fmt.Errorf("error parsing getRecentBlockHashJSONRPCResponse: %w", err)
	}

	return &GetRecentBlockHashResponse{
		Context:   response.Context,
		BlockHash: response.Value.BlockHash,
		FeeCalculator: FeeCalculator{
			blockHash:            response.Value.BlockHash,
			lamportsPerSignature: response.Value.FeeCalculator.LamportsPerSignature,
		},
	}, nil
}

func (j *JSONRPCConnection) SendTransaction(ctx context.Context, request SendTransactionRequest) (*SendTransactionResponse, error) {
	// prepare configuration object
	config := map[string]interface{}{
		"skipPreflight": request.SkipPreflight,
	}

	// prepare transaction as indicated
	var txnData string
	var err error
	switch request.Encoding {
	case Base64Encoding:
		config["encoding"] = Base64Encoding
		txnData, err = request.Transaction.ToBase64()
		if err != nil {
			return nil, fmt.Errorf("error marshalling to base64: %w", err)
		}

	case Base58Encoding:
		fallthrough
	default:
		config["encoding"] = Base58Encoding
		txnData, err = request.Transaction.ToBase58([32]byte{})
		if err != nil {
			return nil, fmt.Errorf("error marshalling to base58: %w", err)
		}
	}

	// set preflightCommitment level if provided
	if request.PreflightCommitmentLevel != "" {
		config["preflightCommitment"] = request.PreflightCommitmentLevel
	}

	// set maxRetries if provided
	if request.MaxRetries != 0 {
		config["maxRetries"] = request.MaxRetries
	}

	// perform rpc call
	rpcResponse, err := j.jsonRPCClient.CallParamArray(
		ctx,
		"sendTransaction",
		nil,
		txnData,
		config,
	)
	if err != nil {
		return nil, fmt.Errorf("error performing sendTransaction json-rpc call: %w", err)
	}
	if rpcResponse.Error != nil {
		return nil, fmt.Errorf("error set on rpc response: %s", rpcResponse.Error.Error())
	}

	// parse response
	response := new(string)
	if err := rpcResponse.GetObject(response); err != nil {
		return nil, fmt.Errorf("error parsing getBalanceJSONRPCResponse: %w", err)
	}

	return &SendTransactionResponse{
		TransactionID: *response,
	}, nil
}
