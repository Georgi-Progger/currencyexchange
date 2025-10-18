package repo

import (
	"context"
	"currencyexchange/internal/models"

	"github.com/jmoiron/sqlx"
)

type currencyRepo struct {
	db *sqlx.DB
}

func NewCurrencyRepo(db *sqlx.DB) *currencyRepo {
	return &currencyRepo{
		db: db,
	}
}

func (c *currencyRepo) GetCurrencies(ctx context.Context) ([]models.Currency, error) {
	query := `
		SELECT id, code, full_name, sign 
		FROM currency
		ORDER BY code;
	`

	currencies := []models.Currency{}
	err := c.db.SelectContext(ctx, &currencies, query)
	if err != nil {
		return nil, err
	}

	return currencies, nil
}

func (c *currencyRepo) GetCurrencyByCode(ctx context.Context, code string) (models.Currency, error) {
	query := `
		SELECT id, code, full_name, sign 
		FROM currency
		WHERE code = $1;
	`

	currency := models.Currency{}
	err := c.db.GetContext(ctx, &currency, query, code)
	if err != nil {
		return models.Currency{}, err
	}

	return currency, nil
}

func (c *currencyRepo) CreateCurrency(ctx context.Context, currency models.Currency) (models.Currency, error) {
	query := `
        INSERT INTO currency (code, full_name, sign) 
        VALUES ($1, $2, $3)
        RETURNING id, code, full_name, sign;
    `

	createdCurrency := models.Currency{}
	err := c.db.QueryRowContext(ctx, query, currency.Code, currency.FullName, currency.Sign).Scan(
		&createdCurrency.Id,
		&createdCurrency.Code,
		&createdCurrency.FullName,
		&createdCurrency.Sign,
	)
	if err != nil {
		return models.Currency{}, err
	}

	return createdCurrency, nil
}

func (c *currencyRepo) GetCurrencyExists(ctx context.Context, code string) (bool, error) {
	query := `
		SELECT EXISTS(SELECT * FROM currency WHERE code = $1);
	`

	var exists bool
	err := c.db.GetContext(ctx, &exists, query, code)
	return exists, err
}
