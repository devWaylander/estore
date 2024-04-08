-- +goose Up
-- +goose StatementBegin

CREATE SCHEMA stocks;

CREATE TABLE IF NOT EXISTS stocks."stock"
(
    id BIGSERIAL PRIMARY KEY,
    sku INT NOT NULL,
    total_count INT NOT NULL,
    reserved INT NOT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA IF EXISTS stocks CASCADE;
-- +goose StatementEnd
