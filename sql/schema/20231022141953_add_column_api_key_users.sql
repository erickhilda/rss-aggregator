-- +goose Up
-- +goose StatementBegin
-- api_key should be unique, and the length should be 64
ALTER TABLE users
ADD COLUMN api_key CHAR(64) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN api_key;

-- +goose StatementEnd