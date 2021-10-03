package jsonrpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type HTTPClient struct {
	urlEndpoint   string
	httpClient    *http.Client
	customHeaders map[string]string
}

func NewHTTPClient(
	urlEndpoint string,
) *HTTPClient {
	return &HTTPClient{
		urlEndpoint: urlEndpoint,
	}
}

// RPCClientOpts can be provided to NewClientWithOpts() to change configuration of HTTPClient.
// HTTPClient: provide a custom http.Client (e.g. to set a proxy, or tls options)
// CustomHeaders: provide custom headers, e.g. to set BasicAuth
type RPCClientOpts struct {
	HTTPClient    *http.Client
	CustomHeaders map[string]string
}

func NewHTTPClientFromOpts(
	endpoint string,
	opts RPCClientOpts,
) *HTTPClient {
	httpClient := NewHTTPClient(
		endpoint,
	)

	if opts.CustomHeaders != nil {
		httpClient.customHeaders = opts.CustomHeaders
	}

	if opts.HTTPClient != nil {
		httpClient.httpClient = opts.HTTPClient
	}

	return httpClient
}

func (c *HTTPClient) CallParamArray(ctx context.Context, method string, additionalHeaders map[string]string, params ...interface{}) (*RPCResponse, error) {
	resp, err := c.call(ctx, method, additionalHeaders, params)
	if err != nil {
		switch typedError := err.(type) {
		case *HTTPError:
			switch typedError.Code {
			case http.StatusBadRequest:
				return nil, fmt.Errorf("%s: %w", typedError.Error(), ErrBadRequest)

			case http.StatusUnauthorized:
				return nil, fmt.Errorf("%s: %w", typedError.Error(), ErrUnauthorized)

			default:
				return nil, fmt.Errorf("%d - %s : %w", typedError.Code, typedError.Error(), ErrHTTPError)
			}

		default:
			// check for connection refused error
			errStr := err.Error()
			if strings.Contains(errStr, "dial") &&
				strings.Contains(errStr, "connection refused") {
				return nil, ErrConnectionRefused
			}

			return nil, err
		}
	}

	if resp == nil {
		return nil, ErrNilResponse
	}

	return resp, nil
}

func (c *HTTPClient) CallParamStruct(ctx context.Context, method string, additionalHeaders map[string]string, params interface{}) (*RPCResponse, error) {
	resp, err := c.call(ctx, method, additionalHeaders, params)
	if err != nil {
		switch typedError := err.(type) {
		case *HTTPError:
			switch typedError.Code {
			case http.StatusBadRequest:
				return nil, fmt.Errorf("%s: %w", typedError.Error(), ErrBadRequest)

			case http.StatusUnauthorized:
				return nil, fmt.Errorf("%s: %w", typedError.Error(), ErrUnauthorized)

			default:
				return nil, fmt.Errorf("%d - %s : %w", typedError.Code, typedError.Error(), ErrHTTPError)
			}

		default:
			// check for connection refused error
			errStr := err.Error()
			if strings.Contains(errStr, "dial") &&
				strings.Contains(errStr, "connection refused") {
				return nil, ErrConnectionRefused
			}

			return nil, err
		}
	}

	if resp == nil {
		return nil, ErrNilResponse
	}

	return resp, nil
}

func (c *HTTPClient) call(ctx context.Context, method string, additionalHeaders map[string]string, params interface{}) (*RPCResponse, error) {
	// construct and marshal rpc request
	rpcRequestData, err := json.Marshal(
		RPCRequest{
			Method:  method,
			Params:  params,
			ID:      1,
			JSONRPC: "2.0",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error json marshalling rpc request: %w", err)
	}

	// construct http request
	httpRequest, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.urlEndpoint,
		bytes.NewReader(rpcRequestData),
	)
	if err != nil {
		return nil, fmt.Errorf("error constructing http request: %w", err)
	}

	// set required headers
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Accept", "application/json")

	// set custom and additional headers
	for k, v := range c.customHeaders {
		httpRequest.Header.Set(k, v)
	}
	for k, v := range additionalHeaders {
		httpRequest.Header.Set(k, v)
	}

	// preform http request
	var httpResponse *http.Response
	if c.httpClient == nil {
		// if http client was not set, use default client
		httpResponse, err = http.DefaultClient.Do(httpRequest)
		if err != nil {
			return nil, fmt.Errorf("rpc call %v() on %v: %v", method, httpRequest.URL.String(), err.Error())
		}
	} else {
		// otherwise http client was set, use it to perform http request
		httpResponse, err = c.httpClient.Do(httpRequest)
		if err != nil {
			return nil, fmt.Errorf("rpc call %v() on %v: %v", method, httpRequest.URL.String(), err.Error())
		}
	}

	// check for an http error
	if httpResponse.StatusCode >= 400 {
		return nil, &HTTPError{
			Code: httpResponse.StatusCode,
			err:  fmt.Errorf("rpc call %v() on %v status code", httpRequest.URL.String(), httpResponse.StatusCode),
		}
	}

	// read body of http response
	httpResponseBodyBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading http resposne body: %w", err)
	}

	// unmarshall http response body to rpc response
	var rpcResponse RPCResponse
	if err := json.Unmarshal(httpResponseBodyBytes, &rpcResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling http response body bytes to rpc response")
	}

	return &rpcResponse, nil
}
