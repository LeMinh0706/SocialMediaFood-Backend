-- +goose Up
-- +goose StatementBegin
CREATE TABLE "reset_password" (
  "id" uuid PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "expires_at" timestamptz NOT NULL,
  "is_active" bool NOT NULL DEFAULT false
);

CREATE INDEX ON "reset_password" ("user_id");

ALTER TABLE "reset_password" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "reset_password" cascade;
-- +goose StatementEnd
