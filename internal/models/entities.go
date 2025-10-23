package models

import (
	"currencyexchange/internal/apperror"

	"github.com/shopspring/decimal"
)

type (
	ExchangeRateRequest struct {
		BaseCurrency   string          `json:"baseCurrency"`
		TargetCurrency string          `json:"targetCurrency"`
		Rate           decimal.Decimal `json:"rate"`
	}
	UpdateExchangeRateRequest struct {
		Rate decimal.Decimal `json:"rate"`
	}
)

func (r *ExchangeRateRequest) Validate() error {
	if r.BaseCurrency == "" || r.TargetCurrency == "" || r.Rate.String() == "" {
		return apperror.ErrValidation
	}

	return nil
}
