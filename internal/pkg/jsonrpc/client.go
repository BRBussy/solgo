package jsonrpc

import (
	"context"
)

type Client interface {
	CallParamArray(ctx context.Context, method string, additionalHeaders map[string]string, params ...interface{}) (*RPCResponse, error)
	CallParamStruct(ctx context.Context, method string, additionalHeaders map[string]string, params interface{}) (*RPCResponse, error)
}
