package usecase

import (
	"context"
	"currencyexchange/internal/apperror"
	"currencyexchange/internal/models"
	"currencyexchange/internal/repo"
	"fmt"
)

type CurrencyUsecase struct {
	repo repo.Currency
}

func NewCurrencyUsecase(repo repo.Currency) *CurrencyUsecase {
	return &CurrencyUsecase{repo: repo}
}

func (c *CurrencyUsecase) GetCurrencies(ctx context.Context) ([]models.Currency, error) {
	modelsList, err := c.repo.GetCurrencies(ctx)
	if err != nil {
		return nil, fmt.Errorf("get currencies is failed: %s", err.Error())
	}

	return modelsList, nil
}

func (c *CurrencyUsecase) GetCurrency(ctx context.Context, code string) (models.Currency, error) {
	return c.repo.GetCurrencyByCode(ctx, code)
}

func (c *CurrencyUsecase) CreateCurrency(ctx context.Context, currency models.Currency) (models.Currency, error) {
	exists, err := c.repo.GetCurrencyExists(ctx, currency.Code)
	if err != nil {
		return models.Currency{}, fmt.Errorf("check currency existence failed: %s", err.Error())
	}
	if exists {
		return models.Currency{}, apperror.ErrCurrencyExists
	}

	model, err := c.repo.CreateCurrency(ctx, currency)
	if err != nil {
		return models.Currency{}, fmt.Errorf("currency create is failed: %s", err.Error())
	}

	return model, nil
}
