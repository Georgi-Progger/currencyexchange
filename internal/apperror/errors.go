package apperror

import "errors"

var (
	ErrCurrencyExists    = errors.New("currency already exists")
	ErrCurrencyNotExists = errors.New("currency not found")
)
