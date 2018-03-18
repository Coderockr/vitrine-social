
-- +migrate Up
CREATE TABLE need_response (
    id           SERIAL PRIMARY KEY,
    need_id     INTEGER NOT NULL REFERENCES needs (id) ON UPDATE CASCADE ON DELETE RESTRICT,
    date         DATE NULL,
    email        VARCHAR(255) NOT NULL,
    name         VARCHAR(255) NOT NULL,
    phone        VARCHAR(255) NOT NULL,
    address      VARCHAR(255) NOT NULL,
    message      VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE need_response;