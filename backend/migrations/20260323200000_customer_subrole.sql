-- +goose Up
ALTER TABLE customer_profiles ADD COLUMN subrole TEXT;

-- +goose Down
ALTER TABLE customer_profiles DROP COLUMN subrole;
