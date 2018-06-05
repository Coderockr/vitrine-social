
-- +migrate Up
ALTER TABLE organizations ALTER COLUMN number TYPE VARCHAR(255);
ALTER TABLE organizations ALTER COLUMN number SET DEFAULT '';

-- +migrate Down
ALTER TABLE organizations ALTER COLUMN number TYPE integer USING number::integer;
ALTER TABLE organizations ALTER COLUMN number SET DEFAULT 0;
