-- +goose Up
ALTER TABLE users ADD COLUMN avatar TEXT;

-- +goose Down
ALTER TABLE users DROP COLUMN avatar;