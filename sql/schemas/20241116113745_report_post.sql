-- +goose Up
-- +goose StatementBegin
CREATE TABLE "issue_post" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "is_deleted" bool NOT NULL DEFAULT false
);
CREATE TABLE "report_post" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "issue_id" int NOT NULL,
  "post_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

INSERT INTO "issue_post" VALUES (1, 'Nội dung không đúng sự thật'),
(2, 'Đạo nhái bài viết'),
(3, 'Bài viết mang nội dung lừa đảo'),
(4, 'Nội dung gây phản cảm'),
(5, 'Có gì đó không ổn');


CREATE INDEX ON "report_post" ("post_id");

CREATE INDEX ON "report_post" ("account_id");

CREATE INDEX ON "report_post" ("issue_id");

ALTER TABLE "report_post" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "report_post" ADD FOREIGN KEY ("issue_id") REFERENCES "issue_post" ("id");

ALTER TABLE "report_post" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "issue_post" cascade;
DROP TABLE IF EXISTS "report_post" cascade;
-- +goose StatementEnd
