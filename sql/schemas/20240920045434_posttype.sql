-- +goose Up
-- +goose StatementBegin
CREATE TABLE "postType" (
  "id" int PRIMARY KEY,
  "name" varchar
);

-- ALTER TABLE "posts" ADD FOREIGN KEY ("post_type_id") REFERENCES "postType" ("id");
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
-- ALTER TABLE "posts" DROP CONSTRAINT IF EXISTS "posts_post_type_id_fkey";
DROP TABLE IF EXISTS "postType"
-- +goose StatementEnd
