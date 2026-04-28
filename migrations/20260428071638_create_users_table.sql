-- +goose Up
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    phone VARCHAR(32),
    email VARCHAR(100) UNIQUE,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    birthday_date DATE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_phone ON users(phone);

-- +goose Down
DROP TABLE users;
