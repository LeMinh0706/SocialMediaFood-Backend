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
  "type" int NOT NULL DEFAULT 3,
  "is_upgrade" bool DEFAULT false,
  "banned" varchar NOT NULL DEFAULT 1,
  "status_id" int DEFAULT 1
);

CREATE TABLE "status" (
  "id" serial PRIMARY KEY,
  "name" varchar
);
INSERT INTO "status" VALUES (1, 'normal');

CREATE INDEX ON "accounts" ("fullname");
CREATE INDEX ON "accounts" ("user_id");
ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "accounts" ADD FOREIGN KEY ("type") REFERENCES "role" ("id");
ALTER TABLE "accounts" ADD FOREIGN KEY ("status_id") REFERENCES "status" ("id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "accounts" CASCADE;
DROP TABLE IF EXISTS "role" CASCADE;
DROP TABLE IF EXISTS "status" CASCADE;
-- +goose StatementEnd
