
-- +migrate Up
ALTER TABLE organizations ALTER COLUMN logo TYPE INTEGER USING (logo::integer);
ALTER TABLE organizations RENAME logo TO logo_image_id;
ALTER TABLE organizations ADD FOREIGN KEY(logo_image_id) REFERENCES organizations_images(id) ON UPDATE CASCADE ON DELETE RESTRICT;
ALTER TABLE organizations ALTER COLUMN logo_image_id DROP NOT NULL;
ALTER TABLE organizations ALTER COLUMN logo_image_id SET DEFAULT NULL;

-- +migrate Down
ALTER TABLE organizations ALTER COLUMN logo_image_id SET NOT NULL;
ALTER TABLE organizations DROP CONSTRAINT "organizations_logo_image_id_fkey";
ALTER TABLE organizations RENAME logo_image_id TO logo;
ALTER TABLE organizations ALTER COLUMN logo TYPE VARCHAR(255);
