
-- +goose Up
CREATE TABLE IF NOT EXISTS tokens (
    token VARCHAR(256) NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    created_a TIMESTAMP NOT NULL DEFAULT NOW(),
    expired TIMESTAMP
);


-- +goose Down
DROP TABLE IF EXISTS tokens;

