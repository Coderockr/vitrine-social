
-- +migrate Up
ALTER TABLE needs ALTER COLUMN updated_at SET DEFAULT now();
-- +migrate Down
ALTER TABLE needs ALTER COLUMN updated_at SET DEFAULT NULL;
