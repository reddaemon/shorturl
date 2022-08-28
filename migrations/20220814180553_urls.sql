-- +goose Up
-- +goose StatementBegin
CREATE TABLE `urls` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `shortUrl` VARCHAR(64) NOT NULL UNIQUE,
    `fullUrl` VARCHAR(64) NOT NULL,
    `created` DATE NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `urls` IF EXISTS;
-- +goose StatementEnd
