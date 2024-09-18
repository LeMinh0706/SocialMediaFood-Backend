-- +goose Up
-- +goose StatementBegin
CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "post_type_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "post_top_id" bigint,
  "description" varchar,
  "date_create_post" bigint NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "posts"
-- +goose StatementEnd
