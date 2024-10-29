-- +goose Up
-- +goose StatementBegin
CREATE TABLE "post_image" (
  "id" bigserial PRIMARY KEY,
  "url_image" varchar NOT NULL,
  "post_id" bigint NOT NULL
);
CREATE INDEX ON "post_image" ("post_id");
ALTER TABLE "post_image" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "post_image" DROP CONSTRAINT IF EXISTS "post_image_post_id_fkey";

DROP TABLE IF EXISTS "post_image";
-- +goose StatementEnd