-- +goose Up
ALTER TABLE brsm_profiles 
ADD COLUMN position VARCHAR(255),
ADD COLUMN phone VARCHAR(50);

ALTER TABLE student_profiles
ADD COLUMN position VARCHAR(255),
ADD COLUMN phone VARCHAR(50),
ADD COLUMN faculty VARCHAR(255),
ADD COLUMN specialty VARCHAR(255),
ADD COLUMN grade FLOAT;

ALTER TABLE university_profiles 
ADD COLUMN position VARCHAR(255),
ADD COLUMN phone VARCHAR(50),
ADD COLUMN department VARCHAR(255);

ALTER TABLE customer_profiles 
ADD COLUMN position VARCHAR(255),
ADD COLUMN phone VARCHAR(50);

-- +goose Down
ALTER TABLE brsm_profiles 
DROP COLUMN position,
DROP COLUMN phone;

ALTER TABLE student_profiles
DROP COLUMN position,
DROP COLUMN phone,
DROP COLUMN faculty,
DROP COLUMN specialty,
DROP COLUMN grade;

ALTER TABLE university_profiles 
DROP COLUMN position,
DROP COLUMN phone,
DROP COLUMN department;

ALTER TABLE customer_profiles 
DROP COLUMN position,
DROP COLUMN phone;