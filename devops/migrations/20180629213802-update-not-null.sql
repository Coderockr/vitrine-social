
-- +migrate Up
UPDATE needs SET updated_at = created_at WHERE updated_at IS NULL;
ALTER TABLE needs ALTER COLUMN updated_at SET NOT NULL;

-- +migrate Down
ALTER TABLE needs ALTER COLUMN updated_at SET NULL;
UPDATE needs SET updated_at = null WHERE updated_at = created_at;
