-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE,
  "hash_pashword" varchar NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "fullname" varchar NOT NULL,
  "gender" int NOT NULL,
  "country" varchar,
  "language" varchar,
  "url_avatar" varchar,
  "role_id" int NOT NULL DEFAULT 1,
  "url_background_profile" varchar,
  "date_create_account" bigint NOT NULL
);

CREATE TABLE "role" (
  "id" int PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "post_image" (
  "id" bigserial PRIMARY KEY,
  "url_image" varchar NOT NULL,
  "post_id" bigint NOT NULL
);

CREATE TABLE "react_post" (
  "id" bigserial PRIMARY KEY,
  "post_id" bigint,
  "user_id" bigint
);

CREATE TABLE "follower" (
  "follow_request" bigint,
  "follow_accept" bigint
);

CREATE INDEX ON "users" ("fullname");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "post_image" ("post_id");

CREATE INDEX ON "react_post" ("post_id");

CREATE INDEX ON "react_post" ("user_id");

CREATE UNIQUE INDEX ON "react_post" ("post_id", "user_id");

CREATE INDEX ON "follower" ("follow_request");

CREATE INDEX ON "follower" ("follow_accept");

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "post_image" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "react_post" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "react_post" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "follower" ADD FOREIGN KEY ("follow_request") REFERENCES "users" ("id");

ALTER TABLE "follower" ADD FOREIGN KEY ("follow_accept") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

INSERT INTO "role" (id, name) VALUES (1, 'Admin'),(2, 'Vip'),(3, 'User');
INSERT INTO "post_type" VALUES (1, 'Normal'),(2, 'Comment');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users" DROP CONSTRAINT IF EXISTS "users_role_id_fkey";
ALTER TABLE "post_image" DROP CONSTRAINT IF EXISTS "post_image_post_id_fkey";
ALTER TABLE "react_post" DROP CONSTRAINT IF EXISTS "react_post_post_id_fkey";
ALTER TABLE "react_post" DROP CONSTRAINT IF EXISTS "react_post_user_id_fkey";
ALTER TABLE "follower" DROP CONSTRAINT IF EXISTS "follower_follow_request_fkey";
ALTER TABLE "follower" DROP CONSTRAINT IF EXISTS "follower_follow_accept_fkey";
ALTER TABLE "posts" DROP CONSTRAINT IF EXISTS "posts_user_id_fkey";

DELETE FROM "role" WHERE id IN (1,2,3);

DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "role";
DROP TABLE IF EXISTS "post_image";
DROP TABLE IF EXISTS "react_post";
DROP TABLE IF EXISTS "follower";


-- +goose StatementEnd
