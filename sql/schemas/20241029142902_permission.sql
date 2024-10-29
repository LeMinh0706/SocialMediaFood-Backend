-- +goose Up
-- +goose StatementBegin
CREATE TABLE "account_permission" (
  "id" bigserial PRIMARY KEY,
  "permission_name" varchar NOT NULL,
  "account_id" bigint NOT NULL,
  "can_select_all" bool NOT NULL DEFAULT false,
  "can_select" bool NOT NULL DEFAULT false,
  "can_insert" bool NOT NULL DEFAULT false,
  "can_update" bool NOT NULL DEFAULT false,
  "can_delete" bool NOT NULL DEFAULT false,
  "can_do_all" bool NOT NULL DEFAULT false
);

CREATE TABLE "permission" (
  "name" varchar PRIMARY KEY
);

CREATE INDEX ON "account_permission" ("account_id");

CREATE INDEX ON "account_permission" ("permission_name");
ALTER TABLE "account_permission" ADD FOREIGN KEY ("permission_name") REFERENCES "permission" ("name");

ALTER TABLE "account_permission" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "account_permission" DROP CONSTRAINT IF EXISTS "account_permission_acccount_id_fkey";
ALTER TABLE "account_permission" DROP CONSTRAINT IF EXISTS "account_permission_permission_name_fkey";
DROP TABLE IF EXISTS "account_permission";
DROP TABLE IF EXISTS "permission";
-- +goose StatementEnd
