package repo

import (
	"context"
	"currencyexchange/internal/models"

	"github.com/jmoiron/sqlx"
)

type Currency interface {
	GetCurrencies(ctx context.Context) ([]models.Currency, error)
	GetCurrencyByCode(ctx context.Context, code string) (models.Currency, error)
	CreateCurrency(ctx context.Context, currency models.Currency) (models.Currency, error)
	GetCurrencyExists(ctx context.Context, code string) (bool, error)
}

type ExchangeRate interface {
	GetExchangeRates(ctx context.Context) ([]models.ExchangeRate, error)
	GetExchangeRateByCode(ctx context.Context, firstCode, secondCode string) (models.ExchangeRate, error)
	CreateExchangeRate(ctx context.Context, rate, firstCode, secondCode string) (models.ExchangeRate, error)
	CheckCurrenciesExist(ctx context.Context, firstCode, secondCode string) (bool, error)
	UpdateExchangeRate(ctx context.Context, firstCode, secondCode, newRate string) (models.ExchangeRate, error)
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
