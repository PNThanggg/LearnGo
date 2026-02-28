-- +goose Up
create table if not exists todos
(
    id         UUID PRIMARY KEY,
    title      VARCHAR(255) NOT NULL,
    completed  BOOLEAN   DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
