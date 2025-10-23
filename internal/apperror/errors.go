package apperror

import (
	"errors"
	"net/http"
)

var (
	ErrCurrencyExists    = errors.New("currency already exists")
	ErrCurrencyNotExists = errors.New("currency not exists")
	ErrValidation        = errors.New("validation error")
	ErrInvalidJSON       = errors.New("invalid json")
	ErrCurrencyNotFound  = errors.New("currency not found")
)

var errorStatusMap = map[error]int{
	ErrCurrencyExists:    http.StatusConflict,
	ErrCurrencyNotExists: http.StatusBadRequest,
	ErrValidation:        http.StatusBadRequest,
	ErrInvalidJSON:       http.StatusBadRequest,
}

func GetHTTPStatus(err error) int {
	if status, ok := errorStatusMap[err]; ok {
		return status
	}
	return http.StatusInternalServerError
}
