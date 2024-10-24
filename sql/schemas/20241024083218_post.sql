-- +goose Up
-- +goose StatementBegin
CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "post_type_id" int NOT NULL,
  "account_id" bigint NOT NULL,
  "post_top_id" bigint,
  "description" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "location" geography(Point,4326),
  "is_banned" bool NOT NULL DEFAULT false,
  "is_deleted" bool NOT NULL DEFAULT false
);
CREATE INDEX ON "posts" ("account_id");

CREATE INDEX ON "posts" ("post_type_id");

CREATE INDEX ON "posts" ("post_top_id");

ALTER TABLE "posts" ADD FOREIGN KEY ("post_type_id") REFERENCES "post_type" ("id");
ALTER TABLE "posts" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
ALTER TABLE "posts" ADD FOREIGN KEY ("post_top_id") REFERENCES "posts" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "posts" DROP CONSTRAINT IF EXISTS "posts_account_id_fkey";
ALTER TABLE "posts" DROP CONSTRAINT IF EXISTS "posts_post_type_id_fkey";
ALTER TABLE "posts" DROP CONSTRAINT IF EXISTS "posts_post_top_id_fkey";

DROP TABLE IF EXISTS "posts";

-- +goose StatementEnd
