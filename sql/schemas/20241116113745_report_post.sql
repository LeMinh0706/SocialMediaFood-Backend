-- +goose Up
-- +goose StatementBegin
CREATE TABLE "issue_post" (
  "id" int PRIMARY KEY,
  "name" varchar NOT NULL,
  "is_deleted" bool NOT NULL DEFAULT false
);
CREATE TABLE "report_post" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "issue_id" int NOT NULL,
  "post_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "report_post" ("post_id");

CREATE INDEX ON "report_post" ("account_id");

CREATE INDEX ON "report_post" ("issue_id");

ALTER TABLE "report_post" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "report_post" ADD FOREIGN KEY ("issue_id") REFERENCES "issue_post" ("id");

ALTER TABLE "report_post" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "issue_post" cascade;
DROP TABLE IF EXISTS "report_post" cascade;
-- +goose StatementEnd
