package models

import "github.com/shopspring/decimal"

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
