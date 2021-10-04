package solana

import "errors"

var (
	ErrUnexpectedNetwork        = errors.New("unexpected network")
	ErrTransactionAlreadySigned = errors.New("transaction already signed")
)
