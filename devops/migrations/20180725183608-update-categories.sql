
-- +migrate Up
UPDATE categories SET name='Higiene Pessoal', slug='higiene-pessoal' WHERE name='Higiene';
INSERT INTO categories (name, slug) VALUES ('Material de Limpeza', 'limpeza');
ALTER TABLE organizations ADD website VARCHAR(255) DEFAULT NULL;

-- +migrate Down
UPDATE categories SET name='Higiene', slug='higiene' WHERE name='Higiene Pessoal';
DELETE FROM categories WHERE slug = 'limpeza';
ALTER TABLE organizations DROP COLUMN website;
