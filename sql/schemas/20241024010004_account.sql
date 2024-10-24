-- +goose Up
-- +goose StatementBegin
CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "fullname" varchar NOT NULL,
  "url_avatar" varchar NOT NULL,
  "url_background_profile" varchar NOT NULL,
  "gender" int,
  "country" varchar,
  "language" varchar,
  "address" varchar,
  "is_deleted" bool NOT NULL DEFAULT false,
  "type" int NOT NULL DEFAULT 3,
  "location" geography(Point,4326),
  "is_upgrade" bool DEFAULT false
);

CREATE INDEX ON "accounts" ("fullname");
ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "accounts_user_id_fkey";

DROP TABLE IF EXISTS "accounts";

-- +goose StatementEnd
