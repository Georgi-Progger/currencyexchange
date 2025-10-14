-- +goose Up
-- +goose StatementBegin
ALTER TABLE exchange_rate ALTER COLUMN rate TYPE DECIMAL(12, 6);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
