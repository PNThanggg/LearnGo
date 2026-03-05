-- migrate:up
CREATE TABLE IF NOT EXISTS refresh_tokens
(
    id            UUID PRIMARY KEY,
    user_id       UUID      NOT NULL,
    refresh_token TEXT      NOT NULL,
    expired_at    TIMESTAMP NOT NULL,
    created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_id_refresh_tokens FOREIGN KEY (user_id) REFERENCES users (id)
);

-- migrate:down
DROP TABLE IF EXISTS refresh_tokens;
