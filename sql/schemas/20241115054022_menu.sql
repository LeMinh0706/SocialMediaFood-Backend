-- +goose Up
-- +goose StatementBegin
CREATE TABLE "menu" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint,
  "dish_name" varchar,
  "quantity" int,
  "price" numeric(11, 2),
  "img" varchar,
  "is_delete" bool NOT NULL DEFAULT false
);
CREATE INDEX ON "menu" ("account_id");
ALTER TABLE "menu" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "menu" cascade;
-- +goose StatementEnd
