
-- +migrate Up
ALTER TABLE needs
  ADD created_at TIMESTAMP NOT NULL DEFAULT now(),
  ADD updated_at TIMESTAMP DEFAULT NULL;

UPDATE categories SET icon = 'roupas' WHERE id = 1;
UPDATE categories SET icon = 'brinquedos' WHERE id = 2;
UPDATE categories SET icon = 'alimentos' WHERE id = 3;
UPDATE categories SET icon = 'servicos' WHERE id = 4;
UPDATE categories SET icon = 'voluntarios' WHERE id = 5;
UPDATE categories SET icon = 'construcao' WHERE id = 6;
UPDATE categories SET icon = 'equipamentos' WHERE id = 7;
UPDATE categories SET icon = 'saude' WHERE id = 8;

-- +migrate Down
ALTER TABLE needs
  DROP COLUMN created_at,
  DROP COLUMN updated_at;

UPDATE categories SET icon = 'http://placehold.it/10x10' WHERE id > 0;
