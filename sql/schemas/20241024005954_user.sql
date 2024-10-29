-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE,
  "username" varchar UNIQUE NOT NULL,
  "hash_password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "is_deleted" bool NOT NULL DEFAULT false
);
CREATE INDEX ON "users" ("username");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
