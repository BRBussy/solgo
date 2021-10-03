package solana

import (
	"context"
	"github.com/BRBussy/solgo/internal/pkg/jsonrpc"
	"reflect"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJSONRPCConnection(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJSONRPCConnection() = %v, want %v", got, tt.want)
			}
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JSONRPCConnection{
				jsonRPCClient: tt.fields.jsonRPCClient,
				config:        tt.fields.config,
			}
			if got := j.Commitment(); got != tt.want {
				t.Errorf("Commitment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONRPCConnection_GetBalance(t *testing.T) {
	type fields struct {
		jsonRPCClient jsonrpc.Client
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &JSONRPCConnection{
				jsonRPCClient: tt.fields.jsonRPCClient,
				config:        tt.fields.config,
			}
			got, err := j.GetBalance(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBalance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithCommitmentLevel(t *testing.T) {
	type args struct {
		c CommitmentLevel
	}
	tests := []struct {
		name string
		args args
		want JSONRPCConnectionOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithCommitmentLevel(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithCommitmentLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithEndpoint(t *testing.T) {
	type args struct {
		e string
	}
	tests := []struct {
		name string
		args args
		want JSONRPCConnectionOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithEndpoint(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithNetwork(t *testing.T) {
	type args struct {
		c Network
	}
	tests := []struct {
		name string
		args args
		want JSONRPCConnectionOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithNetwork(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithNetwork() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonrpcConnectionOptionFunc_apply(t *testing.T) {
	type args struct {
		cfg *jsonrpcConnectionConfig
	}
	tests := []struct {
		name string
		fn   jsonrpcConnectionOptionFunc
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
