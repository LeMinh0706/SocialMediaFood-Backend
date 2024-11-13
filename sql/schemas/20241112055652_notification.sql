-- +goose Up
-- +goose StatementBegin
CREATE TABLE "notification_type" (
  "id" int PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "notification" (
  "id" bigserial PRIMARY KEY,
  "message" varchar NOT NULL,
  "account_id" bigint NOT NULL,
  "type_id" int NOT NULL,
  "post_id" bigint,
  "user_action_id" bigint NOT NULL,
  "invoice_id" bigint,
  "is_seen" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "notification" ("account_id");

CREATE INDEX ON "notification" ("post_id");

CREATE INDEX ON "notification" ("invoice_id");

CREATE INDEX ON "notification" ("type_id");

ALTER TABLE "notification" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "notification" ADD FOREIGN KEY ("type_id") REFERENCES "notification_type" ("id");

ALTER TABLE "notification" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "notification" ADD FOREIGN KEY ("user_action_id") REFERENCES "accounts" ("id");

INSERT INTO "notification_type" VALUES (1, 'For Comment'),(2, 'For like'),(3, 'For follow'),(4, 'For accpet'),(5, 'For invoice'),(6,'For banned');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "notification" cascade;
DROP TABLE IF EXISTS "notification_type" cascade;
-- +goose StatementEnd
