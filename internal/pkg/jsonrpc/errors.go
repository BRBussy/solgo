package jsonrpc

import "errors"

var (
	ErrNilResponse       = errors.New("nil response")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrBadRequest        = errors.New("bad request")
	ErrHTTPError         = errors.New("http error")
	ErrConnectionRefused = errors.New("connection refused")
)
