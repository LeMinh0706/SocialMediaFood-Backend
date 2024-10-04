-- +goose Up
-- +goose StatementBegin
CREATE TABLE "post_type" (
  "id" int PRIMARY KEY,
  "name" varchar
);

-- ALTER TABLE "posts" ADD FOREIGN KEY ("post_type_id") REFERENCES "postType" ("id");
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
-- ALTER TABLE "posts" DROP CONSTRAINT IF EXISTS "posts_post_type_id_fkey";
DELETE FROM "post_type" WHERE id IN (1,2,3);
DROP TABLE IF EXISTS "post_type"
-- +goose StatementEnd
