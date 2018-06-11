
-- +migrate Up
ALTER TABLE categories RENAME icon TO slug;

-- +migrate Down
ALTER TABLE categories RENAME slug TO icon;
