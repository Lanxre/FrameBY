-- +goose Up

CREATE TABLE IF NOT EXISTS squad_statuses (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT
);

INSERT INTO squad_statuses (name, description) VALUES 
    ('pending', 'Ожидает подтверждения'),
    ('recruitment_open', 'Идёт набор'),
    ('closed', 'Закрыта')
ON CONFLICT (name) DO NOTHING;

CREATE TABLE IF NOT EXISTS student_squads (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organizer_id UUID NOT NULL REFERENCES users(id),
    title TEXT NOT NULL,
    description TEXT,
    profile TEXT,
    max_participants INTEGER NOT NULL DEFAULT 1 CHECK (max_participants > 0),
    status_id INTEGER NOT NULL DEFAULT 1 REFERENCES squad_statuses(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_student_squads_organizer ON student_squads(organizer_id);
CREATE INDEX idx_student_squads_status ON student_squads(status_id);

CREATE TABLE IF NOT EXISTS student_squad_participants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    squad_id UUID NOT NULL REFERENCES student_squads(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(squad_id, user_id)
);

CREATE INDEX idx_squad_participants_squad ON student_squad_participants(squad_id);
CREATE INDEX idx_squad_participants_user ON student_squad_participants(user_id);

-- +goose Down
DROP TABLE IF EXISTS student_squad_participants;
DROP TABLE IF EXISTS student_squads;
DROP TABLE IF EXISTS squad_statuses;
