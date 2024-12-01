-- name: IsAdmin :one
SELECT role_id FROM accounts
WHERE user_id = $1
ORDER BY id ASC
LIMIT 1;

-- name: AddUpgradePrice :one
INSERT INTO upgrade_price (
    price 
)VALUES ($1)
RETURNING *;

-- name: GetListUpgradePrice :many
SELECT * FROM upgrade_price
LIMIT $1
OFFSET $2;

-- name: GetLastPrice :one
SELECT * FROM upgrade_price
ORDER BY id DESC
LIMIT 1;

-- name: GetUpgradeQueue :many
SELECT account_id FROM upgrade_queue
WHERE state = 'pending'
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: UpgradeStateQueue :exec
UPDATE upgrade_queue SET state = 'paid'
WHERE account_id = $1;

-- name: UpgradeOwner :exec
UPDATE accounts SET is_upgrade = TRUE AND role_id = 2
WHERE id = $1;
