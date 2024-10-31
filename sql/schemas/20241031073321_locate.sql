-- +goose Up
-- +goose StatementBegin
CREATE TABLE "locate" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "location" geography(Point,4326)
);
ALTER TABLE "locate" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "locate" CASCADE;
-- +goose StatementEnd
