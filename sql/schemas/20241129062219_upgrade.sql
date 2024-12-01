-- +goose Up
-- +goose StatementBegin
CREATE TABLE "upgrade_queue" (
  "account_id" bigint UNIQUE NOT NULL,
  "upgrade_price_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "status" varchar NOT NULL DEFAULT 'pending'
);

CREATE TABLE "upgrade_price" (
  "id" serial PRIMARY KEY,
  "price" DECIMAL(12, 3) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
CREATE INDEX ON "upgrade_queue" ("account_id");

ALTER TABLE "upgrade_queue" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "upgrade_queue" ADD FOREIGN KEY ("upgrade_price_id") REFERENCES "upgrade_price" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "upgrade_queue" cascade;
DROP TABLE IF EXISTS "upgrade_price" cascade;
-- +goose StatementEnd
