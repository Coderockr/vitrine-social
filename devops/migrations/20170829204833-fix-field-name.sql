-- +migrate Up
ALTER TABLE needs RENAME categoty_id TO category_id;
-- +migrate Down
ALTER TABLE needs RENAME category_id TO categoty_id;