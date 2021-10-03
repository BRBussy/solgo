package jsonrpc

// RPCRequest represents a JSON-RPC request object.
// See: http://www.jsonrpc.org/specification#request_object
type RPCRequest struct {
	// Method: string containing the method to be invoked
	Method string `json:"method"`

	// Params: can be nil. if not must be an json array or object
	Params interface{} `json:"params,omitempty"`

	// ID: may always set to 1 for single requests.
	// Should be unique for every request in a batch request.
	ID int `json:"id"`

	// JSONRPC: must always be set to "2.0" for JSON-RPC version 2.0
	JSONRPC string `json:"jsonrpc"`
}
