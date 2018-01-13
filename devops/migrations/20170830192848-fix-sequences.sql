
-- +migrate Up
select setval('organizations_id_seq', 2);
select setval('needs_id_seq', 3);
select setval('categories_id_seq', 9);
-- +migrate Down
