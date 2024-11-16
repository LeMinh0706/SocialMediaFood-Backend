-- +goose Up
-- +goose StatementBegin
CREATE TABLE "react_post" (
  "account_id" bigint NOT NULL,
  "post_id" bigint NOT NULL,
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

DROP TABLE IF EXISTS "react_post" CASCADE;
-- +goose StatementEnd
