-- +goose Up
-- +goose StatementBegin
CREATE TABLE "account_type" (
  "id" int PRIMARY KEY,
  "name" varchar NOT NULL
);

INSERT INTO "account_type" VALUES (1, 'Admin'),(2, 'Owner'),(3, 'User'),(4, 'Reviewer'),(5, 'Moderator');

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
  "is_upgrade" bool DEFAULT false,
  "banned" varchar DEFAULT 0
);

CREATE INDEX ON "accounts" ("fullname");
CREATE INDEX ON "accounts" ("user_id");
ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "accounts" ADD FOREIGN KEY ("type") REFERENCES "account_type" ("id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "accounts_account_type_fkey";
ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "accounts_user_id_fkey";

DROP TABLE IF EXISTS "accounts";
DROP TABLE IF EXISTS "account_type";

-- +goose StatementEnd
