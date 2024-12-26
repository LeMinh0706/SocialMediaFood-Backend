-- +goose Up
-- +goose StatementBegin
CREATE TABLE "rating" (
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "star" int NOT NULL,
  "content" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("from_account_id", "to_account_id")
);
CREATE INDEX ON "rating" ("from_account_id");

CREATE INDEX ON "rating" ("to_account_id");

ALTER TABLE "rating" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "rating" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "rating" cascade;
-- +goose StatementEnd
