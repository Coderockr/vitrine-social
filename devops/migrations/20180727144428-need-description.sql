
-- +migrate Up
ALTER TABLE needs ALTER COLUMN "description" DROP NOT NULL;

-- +migrate Down
ALTER TABLE needs ALTER COLUMN "description" SET NOT NULL;
