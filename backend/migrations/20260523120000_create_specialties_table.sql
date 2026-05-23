-- +goose Up

CREATE TABLE IF NOT EXISTS specialties (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    department_id UUID NOT NULL REFERENCES departments(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (name, department_id)
);

-- Insert existing specialties from student_profiles
INSERT INTO specialties (name, department_id)
SELECT DISTINCT
    TRIM(sp.specialty),
    ud.department_id
FROM student_profiles sp
JOIN university_departments ud ON ud.id = sp.university_department_id
WHERE sp.specialty IS NOT NULL AND TRIM(sp.specialty) != ''
ON CONFLICT (name, department_id) DO NOTHING;

-- Add specialty_id column to student_profiles
ALTER TABLE student_profiles
ADD COLUMN specialty_id UUID REFERENCES specialties(id);

-- Migrate existing specialty references
UPDATE student_profiles sp
SET specialty_id = s.id
FROM specialties s
JOIN university_departments ud ON ud.department_id = s.department_id
WHERE ud.id = sp.university_department_id
  AND s.name = TRIM(sp.specialty);

-- Drop the old free-text specialty column
ALTER TABLE student_profiles
DROP COLUMN specialty;

-- +goose Down
ALTER TABLE student_profiles
ADD COLUMN specialty VARCHAR(255);

UPDATE student_profiles sp
SET specialty = s.name
FROM specialties s
WHERE s.id = sp.specialty_id;

ALTER TABLE student_profiles
DROP COLUMN specialty_id;

DROP TABLE IF EXISTS specialties;
