package exerror

import "errors"

var (
	ErrCurrencyExists = errors.New("currency already exists")
)
