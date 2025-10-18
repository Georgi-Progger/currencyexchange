package repo

import (
	"context"
	"currencyexchange/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type exchangeRateRepo struct {
	db *sqlx.DB
}

func NewExchangeRateRepo(db *sqlx.DB) *exchangeRateRepo {
	return &exchangeRateRepo{
		db: db,
	}
}

func (e *exchangeRateRepo) GetExchangeRates(ctx context.Context) ([]models.ExchangeRate, error) {
	query := `
        SELECT 
            er.id,
            er.rate,
            bc.id as base_id, bc.code as base_code, bc.full_name as base_name, bc.sign as base_sign,
            tc.id as target_id, tc.code as target_code, tc.full_name as target_name, tc.sign as target_sign
        FROM exchange_rate er
        JOIN currency bc ON er.base_currency_id = bc.id
        JOIN currency tc ON er.target_currency_id = tc.id;
    `

	var results []struct {
		ID         int64           `db:"id"`
		Rate       decimal.Decimal `db:"rate"`
		BaseID     int64           `db:"base_id"`
		BaseCode   string          `db:"base_code"`
		BaseName   string          `db:"base_name"`
		BaseSign   string          `db:"base_sign"`
		TargetID   int64           `db:"target_id"`
		TargetCode string          `db:"target_code"`
		TargetName string          `db:"target_name"`
		TargetSign string          `db:"target_sign"`
	}

	err := e.db.SelectContext(ctx, &results, query)
	if err != nil {
		return nil, err
	}

	exchangeRates := make([]models.ExchangeRate, len(results))
	for i, result := range results {
		exchangeRates[i] = models.ExchangeRate{
			Id:   result.ID,
			Rate: result.Rate,
			BaseCurrency: models.Currency{
				Id:       result.BaseID,
				Code:     result.BaseCode,
				FullName: result.BaseName,
				Sign:     result.BaseSign,
			},
			TargetCurrency: models.Currency{
				Id:       result.TargetID,
				Code:     result.TargetCode,
				FullName: result.TargetName,
				Sign:     result.TargetSign,
			},
		}
	}

	return exchangeRates, nil
}

func (e *exchangeRateRepo) GetExchangeRateByCode(ctx context.Context, firstCode, secondCode string) (models.ExchangeRate, error) {
	query := `
		SELECT 
			er.id,
			er.rate,
			bc.id as base_id, bc.code as base_code, bc.full_name as base_name, bc.sign as base_sign,
			tc.id as target_id, tc.code as target_code, tc.full_name as target_name, tc.sign as target_sign
		FROM exchange_rate er
		JOIN currency bc ON er.base_currency_id = bc.id
		JOIN currency tc ON er.target_currency_id = tc.id
		WHERE bc.code = $1 AND tc.code = $2;
	`

	var result struct {
		ID         int64           `db:"id"`
		Rate       decimal.Decimal `db:"rate"`
		BaseID     int64           `db:"base_id"`
		BaseCode   string          `db:"base_code"`
		BaseName   string          `db:"base_name"`
		BaseSign   string          `db:"base_sign"`
		TargetID   int64           `db:"target_id"`
		TargetCode string          `db:"target_code"`
		TargetName string          `db:"target_name"`
		TargetSign string          `db:"target_sign"`
	}

	err := e.db.GetContext(ctx, &result, query, firstCode, secondCode)
	if err != nil {
		return models.ExchangeRate{}, err
	}

	return models.ExchangeRate{
		Id: result.ID,
		BaseCurrency: models.Currency{
			Id:       result.BaseID,
			Code:     result.BaseCode,
			FullName: result.BaseName,
			Sign:     result.BaseSign,
		},
		TargetCurrency: models.Currency{
			Id:       result.TargetID,
			Code:     result.TargetCode,
			FullName: result.TargetName,
			Sign:     result.TargetSign,
		},
		Rate: result.Rate,
	}, nil
}

func (e *exchangeRateRepo) UpdateExchangeRate(ctx context.Context, firstCode, secondCode string, newRate decimal.Decimal) (models.ExchangeRate, error) {
	query := `
        WITH updated AS (
            UPDATE exchange_rate 
            SET rate = $1
            FROM currency bc, currency tc
            WHERE exchange_rate.base_currency_id = bc.id 
            AND exchange_rate.target_currency_id = tc.id
            AND bc.code = $2 
            AND tc.code = $3
            RETURNING exchange_rate.id, exchange_rate.rate, 
                     bc.id as base_id, bc.code as base_code, bc.full_name as base_name, bc.sign as base_sign,
                     tc.id as target_id, tc.code as target_code, tc.full_name as target_name, tc.sign as target_sign
        )
        SELECT * FROM updated;
    `

	var result struct {
		ID         int64           `db:"id"`
		Rate       decimal.Decimal `db:"rate"`
		BaseID     int64           `db:"base_id"`
		BaseCode   string          `db:"base_code"`
		BaseName   string          `db:"base_name"`
		BaseSign   string          `db:"base_sign"`
		TargetID   int64           `db:"target_id"`
		TargetCode string          `db:"target_code"`
		TargetName string          `db:"target_name"`
		TargetSign string          `db:"target_sign"`
	}

	err := e.db.GetContext(ctx, &result, query, newRate, firstCode, secondCode)
	if err != nil {
		return models.ExchangeRate{}, err
	}

	return models.ExchangeRate{
		Id:   result.ID,
		Rate: result.Rate,
		BaseCurrency: models.Currency{
			Id:       result.BaseID,
			Code:     result.BaseCode,
			FullName: result.BaseName,
			Sign:     result.BaseSign,
		},
		TargetCurrency: models.Currency{
			Id:       result.TargetID,
			Code:     result.TargetCode,
			FullName: result.TargetName,
			Sign:     result.TargetSign,
		},
	}, nil
}

func (e *exchangeRateRepo) CreateExchangeRate(ctx context.Context, rate, firstCode, secondCode string) (models.ExchangeRate, error) {
	query := `
		 WITH inserted AS (
            insert into exchange_rate (base_currency_id, target_currency_id, rate) 
			select c1.id, c2.id, $1
			from currency c1 
			join currency c2
			on c1.code = $2 and c2.code = $3
			returning *
        )
        SELECT 
            i.id,
            i.rate,
            bc.id as base_id, bc.code as base_code, bc.full_name as base_name, bc.sign as base_sign,
            tc.id as target_id, tc.code as target_code, tc.full_name as target_name, tc.sign as target_sign
        FROM inserted i
        JOIN currency bc ON i.base_currency_id = bc.id
        JOIN currency tc ON i.target_currency_id = tc.id;
	`

	var result struct {
		Id         int64           `db:"id"`
		Rate       decimal.Decimal `db:"rate"`
		BaseID     int64           `db:"base_id"`
		BaseCode   string          `db:"base_code"`
		BaseName   string          `db:"base_name"`
		BaseSign   string          `db:"base_sign"`
		TargetID   int64           `db:"target_id"`
		TargetCode string          `db:"target_code"`
		TargetName string          `db:"target_name"`
		TargetSign string          `db:"target_sign"`
	}

	err := e.db.GetContext(ctx, &result, query, rate, firstCode, secondCode)
	if err != nil {
		return models.ExchangeRate{}, err
	}

	return models.ExchangeRate{
		Id: result.Id,
		BaseCurrency: models.Currency{
			Id:       result.BaseID,
			Code:     result.BaseCode,
			FullName: result.BaseName,
			Sign:     result.BaseSign,
		},
		TargetCurrency: models.Currency{
			Id:       result.TargetID,
			Code:     result.TargetCode,
			FullName: result.TargetName,
			Sign:     result.TargetSign,
		},
		Rate: result.Rate,
	}, nil
}

func (e *exchangeRateRepo) CheckCurrenciesExist(ctx context.Context, firstCode, secondCode string) (bool, error) {
	query := `
        SELECT 
            EXISTS(SELECT 1 FROM currency WHERE code = $1) AND
            EXISTS(SELECT 1 FROM currency WHERE code = $2)
    `
	var exists bool
	err := e.db.GetContext(ctx, &exists, query, firstCode, secondCode)
	if err != nil {
		return false, err
	}
	return exists, nil
}
