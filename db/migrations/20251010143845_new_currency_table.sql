-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS currency(
    id SERIAL PRIMARY KEY,
    code VARCHAR(3) UNIQUE,
    full_name VARCHAR(255),
    sign VARCHAR(1)
);

CREATE TABLE IF NOT EXISTS exchange_rate(
    id SERIAL PRIMARY KEY,
    base_currency_id INTEGER REFERENCES currency(id) ON DELETE RESTRICT,
    target_currency_id INTEGER REFERENCES currency(id) ON DELETE RESTRICT,
    rate DECIMAL(6)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS exchange_rate;
DROP TABLE IF EXISTS currency;
-- +goose StatementEnd
