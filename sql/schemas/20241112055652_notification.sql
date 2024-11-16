-- +goose Up
-- +goose StatementBegin
CREATE TABLE "invoice" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "create_at" timestamptz NOT NULL,
  "status" varchar,
  "voucher_id" bigint,
  "total" numeric(11, 2)
);

CREATE TABLE "voucher" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "expired_at" bigint NOT NULL,
  "is_delete" bool NOT NULL DEFAULT false,
  "min_to_use" numeric(11, 2),
  "discount" numeric(11, 2),
  "max_amount" numeric(11, 2)
);

CREATE TABLE "apply_voucher" (
  "id" bigserial PRIMARY KEY,
  "voucher_id" bigint NOT NULL,
  "owner_id" bigint NOT NULL
);

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

CREATE INDEX ON "invoice" ("from_account_id");

CREATE INDEX ON "invoice" ("to_account_id");

CREATE INDEX ON "invoice" ("voucher_id");

CREATE INDEX ON "notification" ("account_id");

CREATE INDEX ON "notification" ("post_id");

CREATE INDEX ON "notification" ("invoice_id");

CREATE INDEX ON "notification" ("type_id");

CREATE INDEX ON "apply_voucher" ("voucher_id");

CREATE INDEX ON "apply_voucher" ("owner_id");

ALTER TABLE "apply_voucher" ADD FOREIGN KEY ("voucher_id") REFERENCES "voucher" ("id");

ALTER TABLE "apply_voucher" ADD FOREIGN KEY ("owner_id") REFERENCES "accounts" ("id");

ALTER TABLE "notification" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "notification" ADD FOREIGN KEY ("type_id") REFERENCES "notification_type" ("id");

ALTER TABLE "notification" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "notification" ADD FOREIGN KEY ("invoice_id") REFERENCES "invoice" ("id");

ALTER TABLE "notification" ADD FOREIGN KEY ("user_action_id") REFERENCES "accounts" ("id");

ALTER TABLE "invoice" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "invoice" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "invoice" ADD FOREIGN KEY ("voucher_id") REFERENCES "voucher" ("id");

INSERT INTO "notification_type" VALUES (1, 'For Comment'),(2, 'For like'),(3, 'For follow'),(4, 'For accpet'),(5, 'For invoice'),(6,'For banned');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "apply_voucher" cascade;
DROP TABLE IF EXISTS "voucher" cascade;
DROP TABLE IF EXISTS "invoice" cascade;
DROP TABLE IF EXISTS "notification" cascade;
DROP TABLE IF EXISTS "notification_type" cascade;
-- +goose StatementEnd
