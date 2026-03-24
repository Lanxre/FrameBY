-- +goose Up
ALTER TABLE university_profiles ADD COLUMN faculty TEXT;

-- +goose Down
ALTER TABLE university_profiles DROP COLUMN faculty;
