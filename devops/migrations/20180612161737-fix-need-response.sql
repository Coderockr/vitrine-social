
-- +migrate Up
ALTER TABLE need_response RENAME TO needs_responses;
ALTER TABLE needs_responses DROP COLUMN date;
ALTER TABLE needs_responses ADD created_at TIMESTAMP NOT NULL DEFAULT now();

-- +migrate Down
ALTER TABLE needs_responses DROP COLUMN created_at;
ALTER TABLE needs_responses ADD date DATE NOT NULL DEFAULT now();
ALTER TABLE needs_responses RENAME TO need_response;
