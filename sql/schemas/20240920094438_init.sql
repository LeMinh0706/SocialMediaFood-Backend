-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE,
  "hash_pashword" varchar NOT NULL,
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

CREATE TABLE "postImage" (
  "id" bigserial PRIMARY KEY,
  "url_image" varchar,
  "post_id" bigint
);

CREATE TABLE "reactPost" (
  "id" bigserial PRIMARY KEY,
  "post_id" bigint,
  "user_id" bigint
);

CREATE TABLE "follower" (
  "follow_request" bigint,
  "follow_accept" bigint
);

CREATE INDEX ON "users" ("fullname");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "postImage" ("post_id");

CREATE INDEX ON "reactPost" ("post_id");

CREATE INDEX ON "reactPost" ("user_id");

CREATE UNIQUE INDEX ON "reactPost" ("post_id", "user_id");

CREATE INDEX ON "follower" ("follow_request");

CREATE INDEX ON "follower" ("follow_accept");

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "postImage" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "reactPost" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "reactPost" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "follower" ADD FOREIGN KEY ("follow_request") REFERENCES "users" ("id");

ALTER TABLE "follower" ADD FOREIGN KEY ("follow_accept") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users" DROP CONSTRAINT IF EXISTS "users_role_id_fkey";

ALTER TABLE "users" DROP CONSTRAINT IF EXISTS "users_role_id_fkey";
ALTER TABLE "postImage" DROP CONSTRAINT IF EXISTS "postImage_post_id_fkey";
ALTER TABLE "reactPost" DROP CONSTRAINT IF EXISTS "reactPost_post_id_fkey";
ALTER TABLE "reactPost" DROP CONSTRAINT IF EXISTS "reactPost_user_id_fkey";
ALTER TABLE "follower" DROP CONSTRAINT IF EXISTS "follower_follow_request_fkey";
ALTER TABLE "follower" DROP CONSTRAINT IF EXISTS "follower_follow_accept_fkey";
ALTER TABLE "posts" DROP CONSTRAINT IF EXISTS "posts_user_id_fkey";

DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "role";
DROP TABLE IF EXISTS "postImage";
DROP TABLE IF EXISTS "reactPost";
DROP TABLE IF EXISTS "follower";

-- +goose StatementEnd
