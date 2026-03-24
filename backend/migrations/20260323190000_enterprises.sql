-- +goose Up
CREATE TABLE IF NOT EXISTS enterprises (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    address TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO enterprises (id, name, address) VALUES
    ('b52ab55f-7876-4b76-96d5-a8f0ce25e46f', 'ООО Ромашка', 'г. Гомель, ул. Ленина, 10'),
    ('b52ab55f-7876-4b76-96d5-a8f0ce25e43f', 'Газпром', 'г. Минск, ул. Газеты, 5');

-- +goose Down
DROP TABLE IF EXISTS enterprises;
