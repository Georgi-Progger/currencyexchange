package repo

import (
	"context"
	"currencyexchange/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type Currency interface {
	GetCurrencies(ctx context.Context) ([]model.Currency, error)
	GetCurrencyByCode(ctx context.Context, code string) (*model.Currency, error)
	CreateCurrency(ctx context.Context, code, fullName, sign string) (*model.Currency, error)
	GetCurrencyExists(ctx context.Context, code string) (bool, error)
}

type ExchangeRate interface {
	GetExchangeRates(ctx context.Context) ([]model.ExchangeRate, error)
	GetExchangeRateByCode(ctx context.Context, firstCode, secondCode string) (*model.ExchangeRate, error)
	CreateExchangeRate(ctx context.Context, rate, firstCode, secondCode string) (*model.ExchangeRate, error)
	CheckCurrenciesExist(ctx context.Context, firstCode, secondCode string) (bool, error)
	UpdateExchangeRate(ctx context.Context, firstCode, secondCode string, newRate decimal.Decimal) (*model.ExchangeRate, error)
}

type Repository struct {
	Currency
	ExchangeRate
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Currency:     NewCurrencyRepo(db),
		ExchangeRate: NewExchangeRateRepo(db),
	}
}
