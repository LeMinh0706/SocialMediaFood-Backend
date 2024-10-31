-- +goose Up
-- +goose StatementBegin
CREATE TABLE "permissions" (
  "id" serial PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "role_permission" (
  "id" serial PRIMARY KEY,
  "per_id" bigint,
  "role_id" bigint,
  "can_select_all" bool NOT NULL DEFAULT false,
  "can_select" bool NOT NULL DEFAULT false,
  "can_insert" bool NOT NULL DEFAULT false,
  "can_update" bool NOT NULL DEFAULT false,
  "can_delete" bool NOT NULL DEFAULT false,
  "can_do_all" bool NOT NULL DEFAULT false
);
ALTER TABLE "role_permission" ADD FOREIGN KEY ("per_id") REFERENCES "permissions" ("id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "permissions" CASCADE;
DROP TABLE IF EXISTS "role_permission" CASCADE;
-- +goose StatementEnd
