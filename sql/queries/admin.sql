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
ORDER BY id DESC
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

-- name: GetUpgradeSuccess :many
SELECT account_id FROM upgrade_queue
WHERE state = 'paid'
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: GetListPostReport :many
SELECT post_id, count(account_id) FROM report_post
GROUP BY post_id
HAVING count(account_id)>4
ORDER BY created_at DESC 
LIMIT $1
OFFSET $2;

-- name: ReportPostDetail :many
SELECT r.account_id, r.issue_id, r.created_at, i.* FROM report_post r
LEFT JOIN issue_post i
ON r.issue_id = i.id 
WHERE r.post_id = $1  
ORDER BY r.created_at DESC 
LIMIT $2
OFFSET $3;