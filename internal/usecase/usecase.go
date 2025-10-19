package usecase

import (
	"context"
	"currencyexchange/internal/models"
	"currencyexchange/internal/repo"
)

type Currency interface {
	GetCurrencies(ctx context.Context) ([]models.Currency, error)
	GetCurrency(ctx context.Context, code string) (models.Currency, error)
	CreateCurrency(ctx context.Context, currency models.Currency) (models.Currency, error)
}

type ExchangeRate interface {
	CreateExchangeRate(ctx context.Context, rate, firstCode, secondCode string) (models.ExchangeRate, error)
	GetExchangeRates(ctx context.Context) ([]models.ExchangeRate, error)
	GetExchangeRate(ctx context.Context, twoCodes string) (models.ExchangeRate, error)
	UpdateExchangeRate(ctx context.Context, twoCodes string, newRate string) (models.ExchangeRate, error)
	CalculateExchangeRate(ctx context.Context, baseCurrency, targetCurrenct, amount string) (models.CalculateExchangeRate, error)
}

type Usecase struct {
	Currency
	ExchangeRate
}

func NewUsecase(repo *repo.Repository) *Usecase {
	return &Usecase{
		Currency:     NewCurrencyUsecase(repo.Currency),
		ExchangeRate: NewExchangerateUsecase(repo.ExchangeRate),
	}
}
