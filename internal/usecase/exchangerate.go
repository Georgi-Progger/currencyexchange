package usecase

import (
	"context"
	"currencyexchange/internal/model"
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

func (e *ExchangerateUsecase) CreateExchangeRate(ctx context.Context, rate, firstCode, secondCode string) (*model.ExchangeRate, error) {
	exists, err := e.repo.CheckCurrenciesExist(ctx, firstCode, secondCode)
	if err != nil {
		return nil, fmt.Errorf("check currencies existence failed: %w", err)
	}
	if !exists {
		return nil, errors.New("currencies is not exists")
	}

	model, err := e.repo.CreateExchangeRate(ctx, rate, firstCode, secondCode)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("currency exchange is failed: %s", err.Error()))
	}

	return model, nil
}

func (e *ExchangerateUsecase) GetExchangeRates(ctx context.Context) ([]model.ExchangeRate, error) {
	models, err := e.repo.GetExchangeRates(ctx)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("get exchange rates is failed: %s", err.Error()))
	}

	return models, nil
}

func (e *ExchangerateUsecase) GetExchangeRate(ctx context.Context, twoCodes string) (*model.ExchangeRate, error) {
	if len(twoCodes) != 6 {
		return nil, errors.New("not correct format")
	}

	firstCode := twoCodes[:3]
	secondCode := twoCodes[3:]
	model, err := e.repo.GetExchangeRateByCode(ctx, firstCode, secondCode)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("get currency is failed: %s", err.Error()))
	}

	return model, nil
}

func (e *ExchangerateUsecase) UpdateExchangeRate(ctx context.Context, twoCodes, newRate string) (*model.ExchangeRate, error) {
	return nil, nil
}
