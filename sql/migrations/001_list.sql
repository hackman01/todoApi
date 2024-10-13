-- +goose Up

CREATE TABLE list (
    id UUID PRIMARY KEY,
    item TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down

DROP TABLE list;