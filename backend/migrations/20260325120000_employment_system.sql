-- +goose Up

CREATE TABLE IF NOT EXISTS employment_statuses (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT
);

INSERT INTO employment_statuses (name, description) VALUES 
    ('pending', 'Ожидает подтверждения'),
    ('approved', 'Открыт набор'),
    ('closed', 'Закрыта')
ON CONFLICT (name) DO NOTHING;

INSERT INTO employment_participant_statuses (name, description) VALUES 
    ('applied', 'Подана заявка'),
    ('accepted', 'Принята'),
    ('rejected', 'Отклонена'),
    ('contracted', 'Заключён договор')
ON CONFLICT (name) DO NOTHING;

CREATE TABLE IF NOT EXISTS employment_requests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    enterprise_id UUID NOT NULL REFERENCES enterprises(id),
    university_department_id UUID NOT NULL REFERENCES university_departments(id),
    title TEXT NOT NULL,
    description TEXT,
    requirements TEXT,
    salary TEXT,
    schedule TEXT,
    max_participants INTEGER NOT NULL DEFAULT 1 CHECK (max_participants > 0),
    status_id INTEGER NOT NULL DEFAULT 1 REFERENCES employment_statuses(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_employment_requests_enterprise ON employment_requests(enterprise_id);
CREATE INDEX idx_employment_requests_university ON employment_requests(university_department_id);
CREATE INDEX idx_employment_requests_status ON employment_requests(status_id);

CREATE TABLE IF NOT EXISTS employment_participants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    request_id UUID NOT NULL REFERENCES employment_requests(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    status_id INTEGER NOT NULL DEFAULT 1 REFERENCES employment_participant_statuses(id),
    applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    contracted_at TIMESTAMPTZ,
    UNIQUE(request_id, user_id)
);

CREATE INDEX idx_employment_participants_request ON employment_participants(request_id);
CREATE INDEX idx_employment_participants_user ON employment_participants(user_id);

-- +goose Down
DROP TABLE IF EXISTS employment_participants;
DROP TABLE IF EXISTS employment_requests;
DROP TABLE IF EXISTS employment_participant_statuses;
DROP TABLE IF EXISTS employment_statuses;
