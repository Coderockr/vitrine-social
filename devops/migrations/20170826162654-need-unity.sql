
-- +migrate Up
alter table needs add unity varchar(100) not null;
-- +migrate Down
