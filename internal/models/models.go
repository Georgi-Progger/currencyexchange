package models

import (
	"currencyexchange/internal/apperror"

	"github.com/shopspring/decimal"
)

type (
	Currency struct {
		Id       int64  `json:"id"`
		Code     string `json:"code"`
		FullName string `json:"name" db:"full_name"`
		Sign     string `json:"sign"`
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

func (r *Currency) Validate() error {
	if r.Code == "" || r.FullName == "" || r.Sign == "" {
		return apperror.ErrValidation
	}

	return nil
}
