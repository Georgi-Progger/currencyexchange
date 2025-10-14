package usecase

import (
	"context"
	"currencyexchange/internal/model"
	"currencyexchange/internal/repo"
	"errors"
	"fmt"
)

type CurrencyUsecase struct {
	repo repo.Currency
}

func NewCurrencyUsecase(repo repo.Currency) *CurrencyUsecase {
	return &CurrencyUsecase{repo: repo}
}

func (c *CurrencyUsecase) GetCurrencies(ctx context.Context) ([]model.Currency, error) {
	models, err := c.repo.GetCurrencies(ctx)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("get currencies is failed: %s", err.Error()))
	}

	return models, nil
}

func (c *CurrencyUsecase) GetCurrency(ctx context.Context, code string) (*model.Currency, error) {
	model, err := c.repo.GetCurrencyByCode(ctx, code)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("get currency is failed: %s", err.Error()))
	}

	return model, nil
}

func (c *CurrencyUsecase) CreateCurrency(ctx context.Context, code, fullName, sign string) (*model.Currency, error) {
	exists, err := c.repo.GetCurrencyExists(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("check currency existence failed: %w", err)
	}
	if exists {
		return nil, errors.New("currency already exists")
	}

	model, err := c.repo.CreateCurrency(ctx, code, fullName, sign)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("currency create is failed: %s", err.Error()))
	}

	return model, nil
}
