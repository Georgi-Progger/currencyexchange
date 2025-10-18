package models

import "github.com/shopspring/decimal"

type (
	Currency struct {
		Id       int64  `json:"id"`
		Code     string `json:"code"`
		FullName string `json:"name" db:"full_name"`
		Sign     string `json:"sign"`
	}
	ExchangeRateRequest struct {
		BaseCurrency   string          `json:"baseCurrency"`
		TargetCurrency string          `json:"targetCurrency"`
		Rate           decimal.Decimal `json:"rate"`
	}

	ExchangeRate struct {
		Id             int64           `json:"id"`
		BaseCurrency   Currency        `json:"baseCurrency"`
		TargetCurrency Currency        `json:"targetCurrency"`
		Rate           decimal.Decimal `json:"rate"`
	}

	CalculateExchangeRate struct {
		Id              int64           `json:"id"`
		BaseCurrency    Currency        `json:"baseCurrency"`
		TargetCurrency  Currency        `json:"targetCurrency"`
		Rate            decimal.Decimal `json:"rate"`
		Amount          decimal.Decimal `json:"amount"`
		ConvertedAmount decimal.Decimal `json:"convertedAmount"`
	}
)
