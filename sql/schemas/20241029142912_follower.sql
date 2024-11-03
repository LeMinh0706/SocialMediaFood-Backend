-- +goose Up
-- +goose StatementBegin
CREATE TABLE "follower" (
  "id" bigserial PRIMARY KEY,
  "from_follow" bigint NOT NULL,
  "to_follow" bigint NOT NULL,
  "status" varchar NOT NULL DEFAULT 'pending'
);
CREATE INDEX ON "follower" ("to_follow");
CREATE INDEX ON "follower" ("from_follow");
CREATE UNIQUE INDEX ON "follower" ("to_follow", "from_follow");
ALTER TABLE "follower" ADD FOREIGN KEY ("from_follow") REFERENCES "accounts" ("id");
ALTER TABLE "follower" ADD FOREIGN KEY ("to_follow") REFERENCES "accounts" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "follower" CASCADE;
-- +goose StatementEnd
