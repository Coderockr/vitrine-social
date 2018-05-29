
-- +migrate Up
ALTER TABLE organizations DROP COLUMN address;
ALTER TABLE organizations 
    ADD street VARCHAR(255) NOT NULL DEFAULT '',
    ADD number INT NOT NULL DEFAULT 0,
    ADD complement VARCHAR(255) DEFAULT NULL,
    ADD suburb VARCHAR(255) NOT NULL DEFAULT '',
    ADD city VARCHAR(255) NOT NULL DEFAULT '',
    ADD state VARCHAR(2) NOT NULL DEFAULT '',
    ADD zipcode VARCHAR(9) NOT NULL DEFAULT '',
    ADD created_at TIMESTAMP NOT NULL DEFAULT now();

-- +migrate Down
ALTER TABLE organizations
    DROP COLUMN street, 
    DROP COLUMN number,
    DROP COLUMN complement,
    DROP COLUMN suburb,
    DROP COLUMN city,
    DROP COLUMN state,
    DROP COLUMN zipcode,
    DROP COLUMN created_at;
ALTER TABLE organizations ADD address VARCHAR(255) NOT NULL DEFAULT '';
