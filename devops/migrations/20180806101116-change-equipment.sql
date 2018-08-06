
-- +migrate Up
UPDATE categories SET name = 'Eletroeletrônicos', slug = 'eletroeletronicos' WHERE id = 7;
INSERT INTO categories(name, slug) VALUES ('Recreação', 'recreacao');

ALTER TABLE organizations ADD facebook VARCHAR(255) DEFAULT NULL;
ALTER TABLE organizations ADD instagram VARCHAR(255) DEFAULT NULL;
ALTER TABLE organizations ADD whatsapp VARCHAR(255) DEFAULT NULL;

-- +migrate Down
UPDATE categories SET name = 'Equipamentos', slug = 'equipamentos' WHERE id = 7;
DELETE FROM categories WHERE slug = 'recreacao';

ALTER TABLE organizations DROP COLUMN facebook;
ALTER TABLE organizations DROP COLUMN instagram;
ALTER TABLE organizations DROP COLUMN whatsapp;
