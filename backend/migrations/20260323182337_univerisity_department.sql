-- +goose Up
CREATE TABLE IF NOT EXISTS university_departments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    university_name TEXT NOT NULL,
    department_name TEXT NOT NULL,
    address TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO university_departments (id, university_name, department_name, address) VALUES
    ('a52ab55f-7876-4b76-96d5-a8f0ce25e46f', 'ГГТУ им. П.О. Сухого', 'Кафедра информационных технологий', 'г. Гомель, ул. Кирова, 30'),
    ('a52ab55f-7876-4b76-96d5-a8f0ce25e43f', 'ГГУ им. Ф. Скорины', 'Кафедра программирования', 'г. Гомель, ул. Советская, 104');

-- +goose Down
DROP TABLE IF EXISTS university_departments;
