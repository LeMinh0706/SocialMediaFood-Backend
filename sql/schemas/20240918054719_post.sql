-- +goose Up
-- +goose StatementBegin
CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "post_type_id" int NOT NULL,
  "user_id" bigint NOT NULL,
  "post_top_id" bigint,
  "description" varchar,
  "date_create_post" bigint NOT NULL,
  FOREIGN KEY (post_top_id) REFERENCES posts(id)
);
CREATE INDEX ON "posts" ("user_id");

CREATE INDEX ON "posts" ("post_top_id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "posts"
-- +goose StatementEnd
