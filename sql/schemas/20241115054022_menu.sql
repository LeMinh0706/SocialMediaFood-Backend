-- +goose Up
-- +goose StatementBegin
CREATE TABLE "menu" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "dish_name" varchar NOT NULL,
  "quantity" int NOT NULL,
  "price" DECIMAL(12, 3) NOT NULL,
  "img" varchar NOT NULL,
  "is_delete" bool NOT NULL DEFAULT false
);
CREATE INDEX ON "menu" ("account_id");
ALTER TABLE "menu" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "menu" cascade;
-- +goose StatementEnd
