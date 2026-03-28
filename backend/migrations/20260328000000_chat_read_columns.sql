-- +goose Up
ALTER TABLE chat_members ADD COLUMN IF NOT EXISTS last_read_at TIMESTAMPTZ DEFAULT NOW();
ALTER TABLE chat_members ADD COLUMN IF NOT EXISTS unread_count INT DEFAULT 0;
UPDATE chat_members SET unread_count = 0;

-- +goose Down
ALTER TABLE chat_members DROP COLUMN IF EXISTS last_read_at;
ALTER TABLE chat_members DROP COLUMN IF EXISTS unread_count;
