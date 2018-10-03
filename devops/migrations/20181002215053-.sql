
-- +migrate Up
ALTER TABLE organizations
    DROP COLUMN street,
    DROP COLUMN number,
    DROP COLUMN complement,
    DROP COLUMN city,
    DROP COLUMN state,
    DROP COLUMN zipcode,
    DROP COLUMN neighborhood;
-- +migrate Down
