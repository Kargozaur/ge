-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id TEXT PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE UNIQUE INDEX idx_users_email on users(email);

-- +goose StatementEnd

-- +goose Down
DROP TABLE users;
