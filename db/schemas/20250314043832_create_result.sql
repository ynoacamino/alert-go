-- +goose Up
-- +goose StatementBegin
CREATE TABLE result (
    id TEXT PRIMARY KEY,
    status TEXT NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS result;

-- +goose StatementEnd
