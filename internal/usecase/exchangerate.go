package usecase

import (
	"context"
	"currencyexchange/internal/apperror"
	"currencyexchange/internal/models"
	"currencyexchange/internal/repo"
	"errors"
	"fmt"
)

type ExchangerateUsecase struct {
	repo repo.ExchangeRate
}

func NewExchangerateUsecase(repo repo.ExchangeRate) *ExchangerateUsecase {
	return &ExchangerateUsecase{repo: repo}
}

func (e *ExchangerateUsecase) CreateExchangeRate(ctx context.Context, rate, firstCode, secondCode string) (models.ExchangeRate, error) {
	exists, err := e.repo.CheckCurrenciesExist(ctx, firstCode, secondCode)
	if err != nil {
		return models.ExchangeRate{}, fmt.Errorf("check currencies existence failed: %w", err)
	}
	if !exists {
		return models.ExchangeRate{}, apperror.ErrCurrencyNotExists
	}

	model, err := e.repo.CreateExchangeRate(ctx, rate, firstCode, secondCode)
	if err != nil {
		return models.ExchangeRate{}, errors.New(fmt.Sprintf("currency exchange is failed: %s", err.Error()))
	}

	return model, nil
}

func (e *ExchangerateUsecase) GetExchangeRates(ctx context.Context) ([]models.ExchangeRate, error) {
	models, err := e.repo.GetExchangeRates(ctx)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("get exchange rates is failed: %s", err.Error()))
	}

	return models, nil
}

func (e *ExchangerateUsecase) GetExchangeRate(ctx context.Context, twoCodes string) (models.ExchangeRate, error) {
	if len(twoCodes) != 6 {
		return models.ExchangeRate{}, errors.New("not correct format")
	}

	firstCode := twoCodes[:3]
	secondCode := twoCodes[3:]
	model, err := e.repo.GetExchangeRateByCode(ctx, firstCode, secondCode)
	if err != nil {
		return models.ExchangeRate{}, errors.New(fmt.Sprintf("get currency is failed: %s", err.Error()))
	}

	return model, nil
}

func (e *ExchangerateUsecase) UpdateExchangeRate(ctx context.Context, twoCodes, newRate string) (models.ExchangeRate, error) {
	return models.ExchangeRate{}, nil
}
