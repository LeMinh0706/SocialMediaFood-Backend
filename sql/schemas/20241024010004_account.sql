-- +goose Up
-- +goose StatementBegin
CREATE TABLE "role" (
  "id" serial PRIMARY KEY,
  "name" varchar
);

INSERT INTO "role" VALUES (1, 'Admin'),(2, 'Owner'),(3, 'User'),(4, 'Reviewer'),(5, 'Moderator');

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
  "role_id" int NOT NULL DEFAULT 3,
  "is_upgrade" bool,
  "banned" bigint DEFAULT 0,
  "introduce" varchar
);

CREATE INDEX ON "accounts" ("fullname");
CREATE INDEX ON "accounts" ("user_id");
ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "accounts" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "accounts" CASCADE;
DROP TABLE IF EXISTS "role" CASCADE;
-- +goose StatementEnd
