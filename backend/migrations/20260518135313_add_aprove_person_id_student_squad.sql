-- +goose Up
ALTER TABLE student_squads ADD COLUMN approved_by_university_id UUID REFERENCES users(id) ON DELETE SET NULL;

-- +goose Down
ALTER TABLE student_squads DROP COLUMN approved_by_university_id;
