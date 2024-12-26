-- name: IsAdmin :one
SELECT role_id FROM accounts
WHERE user_id = $1
ORDER BY id ASC
LIMIT 1;

-- name: AddUpgradePrice :one
INSERT INTO upgrade_price (
    title,
    price,
    benefit
)VALUES ($1, $2, $3)
RETURNING *;

-- name: GetListUpgradePrice :many
SELECT * FROM upgrade_price
ORDER BY id DESC
LIMIT $1
OFFSET $2;

-- name: GetChoosePrice :one
SELECT id, title,benefit,price,created_at FROM upgrade_price
WHERE is_choose = TRUE
LIMIT 1;

-- name: GetUpgradeQueue :many
SELECT account_id FROM upgrade_queue
WHERE status = 'pending'
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: UpgradeStateQueue :exec
UPDATE upgrade_queue SET status = 'paid'
WHERE account_id = $1;

-- name: UpgradeOwner :exec
UPDATE accounts SET is_upgrade = TRUE, role_id = 2
WHERE id = $1;

-- name: GetUpgradeSuccess :many
SELECT account_id FROM upgrade_queue
WHERE status = 'paid'
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: GetListPostReport :many
SELECT post_id, COUNT(account_id), MAX(created_at)
FROM report_post
GROUP BY post_id
HAVING COUNT(account_id) >= 5
ORDER BY MAX(created_at) DESC
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

-- name: GetStatusQueue :one
SELECT status FROM upgrade_queue
WHERE account_id = $1
LIMIT 1;

-- name: UnableChoose :exec
UPDATE upgrade_price SET is_choose = FALSE;

-- name: PriceChoosing :exec
UPDATE upgrade_price SET is_choose = TRUE 
WHERE id = $1;

-- name: UpgradeReject :exec
DELETE FROM upgrade_queue WHERE account_id = $1;

-- name: BanPost :exec
UPDATE posts SET is_banned = TRUE 
WHERE id = $1;