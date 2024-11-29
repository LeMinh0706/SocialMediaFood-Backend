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

-- name: GetUpgradePrice :many
SELECT * FROM upgrade_price
LIMIT $1
OFFSET $2;

-- name: GetUpgradeQueue :many
SELECT account_id FROM upgrade_queue
WHERE state = 'pending'
LIMIT $1
OFFSET $2;

-- name: UpgradeStateQueue :exec
UPDATE upgrade_queue SET state = 'paid'
WHERE account_id = $1;

-- name: UpgradeOwner :exec
UPDATE accounts SET role_id = 2
WHERE id = $1;