-- +goose Up
-- +goose StatementBegin
ALTER TABLE "posts" ADD FOREIGN KEY ("post_type_id") REFERENCES "post_type" ("id");
ALTER TABLE posts
ADD CONSTRAINT check_post_top_id CHECK (
    (post_top_id IS NOT NULL AND post_type_id = 2) OR
    (post_top_id IS NULL AND post_type_id <> 2)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "posts" DROP CONSTRAINT IF EXISTS "posts_post_type_id_fkey";
ALTER TABLE "posts" DROP CONSTRAINT IF EXISTS "posts_post_top_id_fkey";
-- +goose StatementEnd
