-- +goose Up
-- +goose StatementBegin
CREATE TABLE "react_post" (
  "id" bigserial PRIMARY KEY,
  "post_id" bigint NOT NULL,
  "account_id" bigint NOT NULL,
  "state" int NOT NULL DEFAULT 1
);
CREATE INDEX ON "react_post" ("post_id");

CREATE INDEX ON "react_post" ("account_id");

CREATE UNIQUE INDEX ON "react_post" ("post_id", "account_id");
ALTER TABLE "react_post" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "react_post" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "react_post" DROP CONSTRAINT IF EXISTS "react_post_post_id_fkey";
ALTER TABLE "react_post" DROP CONSTRAINT IF EXISTS "react_post_account_id_fkey";

DROP TABLE IF EXISTS "react_post";
-- +goose StatementEnd
