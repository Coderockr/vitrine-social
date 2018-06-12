
-- +migrate Up
ALTER TABLE needs RENAME unity TO unit;

-- +migrate Down
ALTER TABLE needs RENAME unit TO unity;
