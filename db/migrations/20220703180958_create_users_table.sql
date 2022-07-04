
-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id SERIAL NOT NULL PRIMARY KEY,
    login VARCHAR(16) NOT NULL,
    password_hash VARCHAR(256) NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS users;

