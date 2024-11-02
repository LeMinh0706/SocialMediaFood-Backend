-- +goose Up
-- +goose StatementBegin
CREATE TABLE "account_status" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "status_id" bigint NOT NULL,
  "created_at" bigint NOT NULL
);

CREATE TABLE "status" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL
);
INSERT INTO "status" VALUES (1, 'normal');
CREATE INDEX ON "account_status" ("account_id");
ALTER TABLE "account_status" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
ALTER TABLE "account_status" ADD FOREIGN KEY ("status_id") REFERENCES "status" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "status" CASCADE;
DROP TABLE IF EXISTS "account_status" CASCADE;
-- +goose StatementEnd
