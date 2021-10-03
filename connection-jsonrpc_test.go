package solana

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/BRBussy/solgo/internal/pkg/jsonrpc"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewJSONRPCConnection(t *testing.T) {
	type args struct {
		opts []JSONRPCConnectionOption
	}
	tests := []struct {
		name string
		args args
		want *JSONRPCConnection
	}{
		{
			name: "default config",
			args: args{},
			want: &JSONRPCConnection{
				jsonRPCClient: jsonrpc.NewHTTPClient(MainnetBeta.MustToRPCURL()),
				config: &jsonrpcConnectionConfig{
					network:         MainnetBeta,
					endpoint:        MainnetBeta.MustToRPCURL(),
					commitmentLevel: ConfirmedCommitmentLevel,
				},
			},
		},
		{
			name: "WithCommitmentLevel config",
			args: args{
				opts: []JSONRPCConnectionOption{
					WithCommitmentLevel(MaxCommitmentLevel),
				},
			},
			want: func() *JSONRPCConnection {
				c := NewJSONRPCConnection()
				c.config.commitmentLevel = MaxCommitmentLevel
				return c
			}(),
		},
		{
			name: "WithNetwork config",
			args: args{
				opts: []JSONRPCConnectionOption{
					WithNetwork(Testnet),
				},
			},
			want: func() *JSONRPCConnection {
				c := NewJSONRPCConnection()
				c.config.network = Testnet
				return c
			}(),
		},
		{
			name: "WithEndpoint config",
			args: args{
				opts: []JSONRPCConnectionOption{
					WithEndpoint("https://someEndpoint.com"),
				},
			},
			want: func() *JSONRPCConnection {
				c := NewJSONRPCConnection()
				c.config.endpoint = "https://someEndpoint.com"
				c.jsonRPCClient = jsonrpc.NewHTTPClient("https://someEndpoint.com")
				return c
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, NewJSONRPCConnection(tt.args.opts...))
		})
	}
}

func TestJSONRPCConnection_Commitment(t *testing.T) {
	type fields struct {
		jsonRPCClient jsonrpc.Client
		config        *jsonrpcConnectionConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   CommitmentLevel
	}{
		{
			name: "basic test",
			fields: fields{
				config: &jsonrpcConnectionConfig{
					commitmentLevel: MaxCommitmentLevel,
				},
			},
			want: MaxCommitmentLevel,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JSONRPCConnection{
				jsonRPCClient: tt.fields.jsonRPCClient,
				config:        tt.fields.config,
			}
			require.Equal(t, tt.want, j.Commitment())
		})
	}
}

func TestJSONRPCConnection_GetAccountInfo(t *testing.T) {
	type fields struct {
		jsonRPCClient *jsonrpc.MockClient
		config        *jsonrpcConnectionConfig
	}
	type args struct {
		ctx     context.Context
		request GetAccountInfoRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetAccountInfoResponse
		wantErr bool
	}{
		{
			name:    "",
			fields:  fields{},
			args:    args{},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.jsonRPCClient.T = t
			j := &JSONRPCConnection{
				jsonRPCClient: tt.fields.jsonRPCClient,
				config:        tt.fields.config,
			}
			got, err := j.GetAccountInfo(tt.args.ctx, tt.args.request)
			require.Equalf(t, tt.wantErr, err != nil, "error is nil")
			require.Equalf(t, tt.want, got, "got neq to want")
		})
	}
}

func TestJSONRPCConnection_GetBalance(t *testing.T) {
	testKeyPair, err := NewRandomKeyPair()
	require.Nil(t, err)

	successfulResponse := GetBalanceResponse{
		Context: Context{
			Slot: 123412356234,
		},
		Value: 100,
	}
	successfulResponseJSONResult, err := json.Marshal(
		GetBalanceJSONRPCResponse{
			Context: successfulResponse.Context,
			Value:   successfulResponse.Value,
		},
	)
	require.Nil(t, err)

	type fields struct {
		jsonRPCClient *jsonrpc.MockClient
		config        *jsonrpcConnectionConfig
	}
	type args struct {
		ctx     context.Context
		request GetBalanceRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetBalanceResponse
		wantErr bool
	}{
		{
			name: "error performing json rpc call - commitment config not provided",
			fields: fields{
				jsonRPCClient: &jsonrpc.MockClient{
					CallParamArrayFunc: func(t *testing.T, m *jsonrpc.MockClient, ctx context.Context, method string, additionalHeaders map[string]string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
						require.Equalf(
							t,
							[]interface{}{
								testKeyPair.PublicKey.ToBase58(),
								map[string]interface{}{
									"commitment": MaxCommitmentLevel,
								},
							},
							params,
							"params not as expected",
						)

						return nil, errors.New("some err")
					},
				},
				config: &jsonrpcConnectionConfig{
					commitmentLevel: MaxCommitmentLevel,
				},
			},
			args: args{
				ctx: context.Background(),
				request: GetBalanceRequest{
					PublicKey: testKeyPair.PublicKey,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error set on rpc response - commitment config provided",
			fields: fields{
				jsonRPCClient: &jsonrpc.MockClient{
					CallParamArrayFunc: func(t *testing.T, m *jsonrpc.MockClient, ctx context.Context, method string, additionalHeaders map[string]string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
						require.Equalf(
							t,
							[]interface{}{
								testKeyPair.PublicKey.ToBase58(),
								map[string]interface{}{
									"commitment": ProcessedCommitmentLevel,
								},
							},
							params,
							"params not as expected",
						)

						return &jsonrpc.RPCResponse{
							Error: &jsonrpc.RPCError{Message: "bad things happened"},
						}, nil
					},
				},
				config: &jsonrpcConnectionConfig{
					commitmentLevel: MaxCommitmentLevel,
				},
			},
			args: args{
				ctx: context.Background(),
				request: GetBalanceRequest{
					PublicKey:       testKeyPair.PublicKey,
					CommitmentLevel: ProcessedCommitmentLevel,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error parsing GetBalanceJSONRPCResponse",
			fields: fields{
				jsonRPCClient: &jsonrpc.MockClient{
					CallParamArrayFunc: func(t *testing.T, m *jsonrpc.MockClient, ctx context.Context, method string, additionalHeaders map[string]string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
						return &jsonrpc.RPCResponse{
							Result: []byte("invalid data here"),
							Error:  nil,
						}, nil
					},
				},
				config: &jsonrpcConnectionConfig{
					commitmentLevel: MaxCommitmentLevel,
				},
			},
			args: args{
				ctx: context.Background(),
				request: GetBalanceRequest{
					PublicKey:       testKeyPair.PublicKey,
					CommitmentLevel: FinalizedCommitmentLevel,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				jsonRPCClient: &jsonrpc.MockClient{
					CallParamArrayFunc: func(t *testing.T, m *jsonrpc.MockClient, ctx context.Context, method string, additionalHeaders map[string]string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
						return &jsonrpc.RPCResponse{
							Result: successfulResponseJSONResult,
							Error:  nil,
						}, nil
					},
				},
				config: &jsonrpcConnectionConfig{
					commitmentLevel: MaxCommitmentLevel,
				},
			},
			args: args{
				ctx: context.Background(),
				request: GetBalanceRequest{
					PublicKey:       testKeyPair.PublicKey,
					CommitmentLevel: MaxCommitmentLevel,
				},
			},
			want:    &successfulResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.jsonRPCClient.T = t
			j := &JSONRPCConnection{
				jsonRPCClient: tt.fields.jsonRPCClient,
				config:        tt.fields.config,
			}
			got, err := j.GetBalance(tt.args.ctx, tt.args.request)
			require.Equalf(t, tt.wantErr, err != nil, "error is nil")
			require.Equalf(t, tt.want, got, "got neq to want")
		})
	}
}
