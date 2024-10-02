-- +goose Up
-- +goose StatementBegin
INSERT INTO "role" (id, name) VALUES (1, 'Admin'),(2, 'Vip'),(3, 'User');
INSERT INTO "post_type" VALUES (1, 'Normal'),(2, 'Comment'),(3, 'Prime');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM "role" WHERE id IN (1,2,3);
DELETE FROM "post_type" WHERE id IN (1,2,3);
-- +goose StatementEnd
