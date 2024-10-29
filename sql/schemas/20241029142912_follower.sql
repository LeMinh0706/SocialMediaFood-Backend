-- +goose Up
-- +goose StatementBegin
CREATE TABLE "follower" (
  "id" bigserial PRIMARY KEY,
  "follow_request" bigint NOT NULL,
  "follow_accept" bigint NOT NULL,
  "status" varchar NOT NULL DEFAULT 'pending'
);
CREATE INDEX ON "follower" ("follow_request");
CREATE INDEX ON "follower" ("follow_accept");
CREATE UNIQUE INDEX ON "follower" ("follow_request", "follow_accept");
ALTER TABLE "follower" ADD FOREIGN KEY ("follow_request") REFERENCES "accounts" ("id");
ALTER TABLE "follower" ADD FOREIGN KEY ("follow_accept") REFERENCES "accounts" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "follower" DROP CONSTRAINT IF EXISTS "follow_request_fkey";
ALTER TABLE "follower" DROP CONSTRAINT IF EXISTS "follow_accept_fkey";
DROP TABLE IF EXISTS "follower"
-- +goose StatementEnd
