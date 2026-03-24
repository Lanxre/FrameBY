-- +goose Up

CREATE TABLE IF NOT EXISTS universities (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS departments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

ALTER TABLE university_departments 
ADD COLUMN university_id UUID REFERENCES universities(id),
ADD COLUMN department_id UUID REFERENCES departments(id);

INSERT INTO universities (id, name)
SELECT DISTINCT gen_random_uuid(), university_name 
FROM university_departments;

INSERT INTO departments (id, name)
SELECT DISTINCT gen_random_uuid(), department_name 
FROM university_departments;

UPDATE university_departments ud
SET university_id = u.id
FROM universities u
WHERE u.name = ud.university_name;

UPDATE university_departments ud
SET department_id = d.id
FROM departments d
WHERE d.name = ud.department_name;

ALTER TABLE university_departments 
ALTER COLUMN university_id SET NOT NULL,
ALTER COLUMN department_id SET NOT NULL;

ALTER TABLE university_departments 
DROP COLUMN university_name,
DROP COLUMN department_name;

ALTER TABLE student_profiles 
ADD COLUMN university_department_id UUID REFERENCES university_departments(id);

ALTER TABLE university_profiles 
ADD COLUMN university_department_id UUID REFERENCES university_departments(id);

UPDATE student_profiles sp
SET university_department_id = (
    SELECT ud.id FROM university_departments ud
    JOIN universities u ON ud.university_id = u.id
    JOIN departments d ON ud.department_id = d.id
    WHERE u.name = sp.faculty
    LIMIT 1
);

UPDATE university_profiles up
SET university_department_id = (
    SELECT ud.id FROM university_departments ud
    JOIN universities u ON ud.university_id = u.id
    JOIN departments d ON ud.department_id = d.id
    WHERE u.name = up.faculty AND d.name = up.department
    LIMIT 1
);

ALTER TABLE student_profiles 
DROP COLUMN faculty;

ALTER TABLE university_profiles 
DROP COLUMN faculty,
DROP COLUMN department;

-- +goose Down
ALTER TABLE student_profiles ADD COLUMN faculty TEXT;
ALTER TABLE university_profiles ADD COLUMN faculty TEXT, ADD COLUMN department TEXT;

UPDATE student_profiles sp
SET faculty = (
    SELECT u.name FROM university_departments ud
    JOIN universities u ON ud.university_id = u.id
    WHERE ud.id = sp.university_department_id
    LIMIT 1
);

UPDATE university_profiles up
SET faculty = (
    SELECT u.name FROM university_departments ud
    JOIN universities u ON ud.university_id = u.id
    WHERE ud.id = up.university_department_id
    LIMIT 1
),
department = (
    SELECT d.name FROM university_departments ud
    JOIN departments d ON ud.department_id = d.id
    WHERE ud.id = up.university_department_id
    LIMIT 1
);

ALTER TABLE student_profiles DROP COLUMN university_department_id;
ALTER TABLE university_profiles DROP COLUMN university_department_id;

ALTER TABLE university_departments ADD COLUMN university_name TEXT, ADD COLUMN department_name TEXT;

UPDATE university_departments ud
SET university_name = u.name, department_name = d.name
FROM universities u, departments d
WHERE ud.university_id = u.id AND ud.department_id = d.id;

ALTER TABLE university_departments ALTER COLUMN university_name SET NOT NULL, ALTER COLUMN department_name SET NOT NULL;

DROP TABLE IF EXISTS departments;
DROP TABLE IF EXISTS universities;
