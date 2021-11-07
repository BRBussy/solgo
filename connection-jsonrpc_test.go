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
	testKeyPair, err := NewRandomKeyPair()
	require.Nil(t, err)

	type fields struct {
		jsonRPCClient *jsonrpc.MockClient
		config        *jsonrpcConnectionConfig
	}
	type args struct {
		ctx     context.Context
		request GetAccountInfoRequest
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        *GetAccountInfoResponse
		wantToCheck func(t *testing.T, gotResponse *GetAccountInfoResponse)
		wantErr     bool
	}{
		{
			name: "error performing CallParamArray - neither commitment nor encoding given",
			fields: fields{
				jsonRPCClient: &jsonrpc.MockClient{
					CallParamArrayFunc: func(t *testing.T, m *jsonrpc.MockClient, ctx context.Context, method string, additionalHeaders map[string]string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
						require.Equalf(t, "getAccountInfo", method, "method not as expected")
						require.Equalf(
							t,
							[]interface{}{
								testKeyPair.PublicKey.ToBase58(),
								map[string]interface{}{
									"commitment": ConfirmedCommitmentLevel,
									"encoding":   Base64Encoding,
								},
							},
							params,
							"params not as expected",
						)

						return nil, errors.New("some err")
					},
				},
				config: &jsonrpcConnectionConfig{
					commitmentLevel: ConfirmedCommitmentLevel,
				},
			},
			args: args{
				ctx: nil,
				request: GetAccountInfoRequest{
					PublicKey:       testKeyPair.PublicKey,
					CommitmentLevel: "",
					Encoding:        "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error set on rpc response in CallParamArray - both commitment and encoding given",
			fields: fields{
				jsonRPCClient: &jsonrpc.MockClient{
					CallParamArrayFunc: func(t *testing.T, m *jsonrpc.MockClient, ctx context.Context, method string, additionalHeaders map[string]string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
						require.Equalf(t, "getAccountInfo", method, "method not as expected")
						require.Equalf(
							t,
							[]interface{}{
								testKeyPair.PublicKey.ToBase58(),
								map[string]interface{}{
									"commitment": MaxCommitmentLevel,
									"encoding":   JSONParsedEncoding,
								},
							},
							params,
							"params not as expected",
						)

						return &jsonrpc.RPCResponse{
							Error: &jsonrpc.RPCError{
								Message: "Bad things happened",
							},
						}, nil
					},
				},
				config: &jsonrpcConnectionConfig{
					commitmentLevel: ConfirmedCommitmentLevel,
				},
			},
			args: args{
				ctx: nil,
				request: GetAccountInfoRequest{
					PublicKey:       testKeyPair.PublicKey,
					CommitmentLevel: MaxCommitmentLevel,
					Encoding:        JSONParsedEncoding,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error parsing for JSONParsedEncoding",
			fields: fields{
				jsonRPCClient: &jsonrpc.MockClient{
					CallParamArrayFunc: func(t *testing.T, m *jsonrpc.MockClient, ctx context.Context, method string, additionalHeaders map[string]string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
						return &jsonrpc.RPCResponse{
							Result: json.RawMessage("not gonna parse"),
						}, nil
					},
				},
				config: &jsonrpcConnectionConfig{
					commitmentLevel: ConfirmedCommitmentLevel,
				},
			},
			args: args{
				ctx: nil,
				request: GetAccountInfoRequest{
					PublicKey: testKeyPair.PublicKey,
					Encoding:  JSONParsedEncoding,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error parsing for any other encoding",
			fields: fields{
				jsonRPCClient: &jsonrpc.MockClient{
					CallParamArrayFunc: func(t *testing.T, m *jsonrpc.MockClient, ctx context.Context, method string, additionalHeaders map[string]string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
						return &jsonrpc.RPCResponse{
							Result: json.RawMessage("not gonna parse"),
						}, nil
					},
				},
				config: &jsonrpcConnectionConfig{
					commitmentLevel: ConfirmedCommitmentLevel,
				},
			},
			args: args{
				ctx: nil,
				request: GetAccountInfoRequest{
					PublicKey: testKeyPair.PublicKey,
					Encoding:  Base58Encoding,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success for JSONParsedEncoding",
			fields: fields{
				jsonRPCClient: &jsonrpc.MockClient{
					CallParamArrayFunc: func(t *testing.T, m *jsonrpc.MockClient, ctx context.Context, method string, additionalHeaders map[string]string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
						return &jsonrpc.RPCResponse{
							Result: json.RawMessage(`{
    "context": {
      "slot": 1
    },
    "value": {
      "data": {
        "nonce": {
          "initialized": {
            "authority": "Bbqg1M4YVVfbhEzwA9SpC9FhsaG83YMTYoR4a8oTDLX",
            "blockhash": "3xLP3jK6dVJwpeGeTDYTwdDK3TKchUf1gYYGHa4sF3XJ",
            "feeCalculator": {
              "lamportsPerSignature": 5000
            }
          }
        }
      },
      "executable": false,
      "lamports": 1000000000,
      "owner": "11111111111111111111111111111111",
      "rentEpoch": 2
    }
}`),
						}, nil
					},
				},
				config: &jsonrpcConnectionConfig{
					commitmentLevel: ConfirmedCommitmentLevel,
				},
			},
			args: args{
				ctx: nil,
				request: GetAccountInfoRequest{
					PublicKey: testKeyPair.PublicKey,
					Encoding:  JSONParsedEncoding,
				},
			},
			wantToCheck: func(t *testing.T, gotResponse *GetAccountInfoResponse) {
				require.Equalf(
					t,
					Context{Slot: 1},
					gotResponse.Context,
					"context not as expected",
				)

				// ensure account info type is as expected
				typedAccountInfo, ok := gotResponse.AccountInfo.(AccountInfoJSONData)
				require.Truef(t, ok, "unable to cast to AccountInfoJSONData")

				data := map[string]json.RawMessage{
					"nonce": json.RawMessage(`{
				"initialized": {
				 "authority": "Bbqg1M4YVVfbhEzwA9SpC9FhsaG83YMTYoR4a8oTDLX",
				 "blockhash": "3xLP3jK6dVJwpeGeTDYTwdDK3TKchUf1gYYGHa4sF3XJ",
				 "feeCalculator": {
				   "lamportsPerSignature": 5000
				 }
				}
}`),
				}

				// extract nonce data
				expectedNonceData := data["nonce"]
				gotNonceData, found := typedAccountInfo.Data["nonce"]
				require.Truef(t, found, "nonce data not found on received data")

				require.JSONEqf(
					t,
					string(expectedNonceData),
					string(gotNonceData),
					"json data not as expected",
				)

			},
			wantErr: false,
		},
		{
			name: "success for base58 encoding",
			fields: fields{
				jsonRPCClient: &jsonrpc.MockClient{
					CallParamArrayFunc: func(t *testing.T, m *jsonrpc.MockClient, ctx context.Context, method string, additionalHeaders map[string]string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
						return &jsonrpc.RPCResponse{
							Result: json.RawMessage(`{
    "context": {
      "slot": 1
    },
    "value": {
      "data": [
        "11116bv5nS2h3y12kD1yUKeMZvGcKLSjQgX6BeV7u1FrjeJcKfsHRTPuR3oZ1EioKtYGiYxpxMG5vpbZLsbcBYBEmZZcMKaSoGx9JZeAuWf",
        "base58"
      ],
      "executable": false,
      "lamports": 1000000000,
      "owner": "11111111111111111111111111111111",
      "rentEpoch": 2
    }
  }`),
						}, nil
					},
				},
				config: &jsonrpcConnectionConfig{
					commitmentLevel: ConfirmedCommitmentLevel,
				},
			},
			args: args{
				ctx: nil,
				request: GetAccountInfoRequest{
					PublicKey: testKeyPair.PublicKey,
					Encoding:  Base58Encoding,
				},
			},
			want: &GetAccountInfoResponse{
				Context: Context{Slot: 1},
				AccountInfo: AccountInfoEncodedData{
					Executable: false,
					Lamports:   1000000000,
					Data: []string{
						"11116bv5nS2h3y12kD1yUKeMZvGcKLSjQgX6BeV7u1FrjeJcKfsHRTPuR3oZ1EioKtYGiYxpxMG5vpbZLsbcBYBEmZZcMKaSoGx9JZeAuWf",
						"base58",
					},
					Owner:     "11111111111111111111111111111111",
					RentEpoch: 2,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fields.jsonRPCClient != nil {
				tt.fields.jsonRPCClient.T = t
			}
			j := &JSONRPCConnection{
				jsonRPCClient: tt.fields.jsonRPCClient,
				config:        tt.fields.config,
			}
			got, err := j.GetAccountInfo(tt.args.ctx, tt.args.request)
			require.Equalf(t, tt.wantErr, err != nil, "error is nil")
			if tt.wantToCheck == nil {
				require.Equalf(t, tt.want, got, "got neq to want")
			} else {
				tt.wantToCheck(t, got)
			}
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
		getBalanceJSONRPCResponse{
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
						require.Equalf(t, "getBalance", method, "method not as expected")
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
			name: "error parsing getBalanceJSONRPCResponse",
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
			if tt.fields.jsonRPCClient != nil {
				tt.fields.jsonRPCClient.T = t
			}
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
