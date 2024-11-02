-- +goose Up
-- +goose StatementBegin
CREATE TABLE "post_type" (
  "id" int PRIMARY KEY,
  "name" varchar NOT NULL
);
INSERT INTO "post_type" VALUES (1, 'Normal'),(2, 'Premium'),(3, 'Avatar'),(4, 'Background'),(9, 'Comment');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "post_type";
-- +goose StatementEnd
