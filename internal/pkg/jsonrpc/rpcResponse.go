package jsonrpc

import (
	"encoding/json"
	"strconv"
)

// RPCResponse represents a JSON-RPC response object.
// See: http://www.jsonrpc.org/specification#response_object
type RPCResponse struct {
	// JSONRPC: must always be set to "2.0" for JSON-RPC version 2.0
	JSONRPC string `json:"jsonrpc"`

	// Result: holds the result of the rpc call if no error occurred, nil otherwise.
	Result json.RawMessage `json:"result,omitempty"`

	// Error: holds an RPCError object if an error occurred. must be nil on success.
	Error *RPCError `json:"error,omitempty"`

	// ID: may always be 0 for single requests, should be unique for batch requests.
	ID int `json:"id"`
}

// RPCError represents a JSON-RPC error object if an RPC error occurred.
// See: http://www.jsonrpc.org/specification#error_object
type RPCError struct {
	// Code: holds the error code
	Code int `json:"code"`

	// Message: holds a short error message
	Message string `json:"message"`

	// Data: holds additional error data, may be nil
	Data interface{} `json:"data,omitempty"`
}

// Error function is provided to be used as error object.
func (e *RPCError) Error() string {
	return strconv.Itoa(e.Code) + ":" + e.Message
}

// GetObject converts the rpc response to an arbitrary type.
// The function works as you would expect it from json.Unmarshal()
func (RPCResponse *RPCResponse) GetObject(target interface{}) error {
	return json.Unmarshal(RPCResponse.Result, target)
}

// HTTPError represents a error that occurred on HTTP level.
//
// An error of type HTTPError is returned when a HTTP error occurred (status code)
// and the body could not be parsed to a valid RPCResponse object that holds a RPCError.
//
// Otherwise a RPCResponse object is returned with a RPCError field that is not nil.
type HTTPError struct {
	Code int
	err  error
}

// Error function is provided to be used as error object.
func (e *HTTPError) Error() string {
	return e.err.Error()
}
