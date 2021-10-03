package jsonrpc

import (
	"context"
	"testing"
)

type MockClient struct {
	T                              *testing.T
	CallParamArrayFunc             func(t *testing.T, m *MockClient, ctx context.Context, method string, additionalHeaders map[string]string, params ...interface{}) (*RPCResponse, error)
	CallParamArrayFuncInvocations  int
	CallParamStructFunc            func(t *testing.T, m *MockClient, ctx context.Context, method string, additionalHeaders map[string]string, params interface{}) (*RPCResponse, error)
	CallParamStructFuncInvocations int
}

func (m *MockClient) CallParamArray(ctx context.Context, method string, additionalHeaders map[string]string, params ...interface{}) (*RPCResponse, error) {
	m.CallParamArrayFuncInvocations++
	if m.CallParamArrayFunc == nil {
		return nil, nil
	}
	return m.CallParamArrayFunc(m.T, m, ctx, method, additionalHeaders, params...)
}

func (m *MockClient) CallParamStruct(ctx context.Context, method string, additionalHeaders map[string]string, params interface{}) (*RPCResponse, error) {
	m.CallParamStructFuncInvocations++
	if m.CallParamStructFunc == nil {
		return nil, nil
	}
	return m.CallParamStructFunc(m.T, m, ctx, method, additionalHeaders, params)
}
