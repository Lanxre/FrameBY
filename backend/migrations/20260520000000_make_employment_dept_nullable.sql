-- +goose Up
ALTER TABLE employment_requests ALTER COLUMN university_department_id DROP NOT NULL;

-- +goose Down
ALTER TABLE employment_requests ALTER COLUMN university_department_id SET NOT NULL;
