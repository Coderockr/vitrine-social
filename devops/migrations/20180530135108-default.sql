
-- +migrate Up
ALTER TABLE organizations RENAME resume TO about;
ALTER TABLE organizations RENAME suburb TO neighborhood;

-- +migrate Down
ALTER TABLE organizations RENAME about TO resume;
ALTER TABLE organizations RENAME neighborhood TO suburb;
