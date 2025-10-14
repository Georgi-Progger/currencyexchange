package usecase

import (
	"context"
	"currencyexchange/internal/model"
	"currencyexchange/internal/repo"
)

type Currency interface {
	GetCurrencies(ctx context.Context) ([]model.Currency, error)
	GetCurrency(ctx context.Context, code string) (*model.Currency, error)
	CreateCurrency(ctx context.Context, code, fullName, sign string) (*model.Currency, error)
}

type ExchangeRate interface {
	CreateExchangeRate(ctx context.Context, rate, firstCode, secondCode string) (*model.ExchangeRate, error)
	GetExchangeRates(ctx context.Context) ([]model.ExchangeRate, error)
	GetExchangeRate(ctx context.Context, twoCodes string) (*model.ExchangeRate, error)
	UpdateExchangeRate(ctx context.Context, twoCodes string, newRate string) (*model.ExchangeRate, error)
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
