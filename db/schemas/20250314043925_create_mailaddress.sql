-- +goose Up
-- +goose StatementBegin
CREATE TABLE mailAddress (
    id TEXT PRIMARY KEY,
    address TEXT NOT NULL UNIQUE,
    active BOOLEAN DEFAULT 1 NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS mailAddress;

-- +goose StatementEnd
