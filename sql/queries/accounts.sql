-- name: CreateAccounts :one
INSERT INTO accounts(
    user_id,
    fullname,
    gender,
    country,
    language,
    role_id,
    url_avatar,
    url_background_profile,
    is_upgrade
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, false
) RETURNING * ;

-- name: GetAccountByUserId :many
SELECT id
FROM accounts
WHERE user_id = $1
ORDER BY id;

-- name: GetDetailAccount :one
SELECT * FROM accounts
WHERE id = $1;

-- name: GetAccountById :one
SELECT id, fullname, url_avatar, url_background_profile, role_id FROM accounts
WHERE id = $1
LIMIT 1;

-- name: UpdateName :one
UPDATE accounts SET fullname = $2
WHERE id = $1
RETURNING id, fullname;

-- name: UpdateAvatar :one
UPDATE accounts SET url_avatar = $2
WHERE id = $1
RETURNING id, fullname, url_avatar, url_background_profile, role_id;

-- name: SearchingAccounts :many
SELECT id, fullname, url_avatar, url_background_profile, role_id FROM accounts
WHERE fullname ILIKE '%' || $1 || '%'
LIMIT $2
OFFSET $3;

-- name: UpdateBackground :one
UPDATE accounts SET url_background_profile = $2
WHERE id = $1
RETURNING id, fullname, url_avatar, url_background_profile, role_id;

-- name: UpgradeSuccess :one
UPDATE accounts SET is_upgrade = TRUE
WHERE id = $1
RETURNING id;

-- name: UpdateEmail :exec
UPDATE users SET email = $2
WHERE id = $1;

-- name: UpgradeOnwerRequest :exec
INSERT INTO upgrade_queue (
    account_id,
    upgrade_price_id
)VALUES (
    $1, $2
)RETURNING *;

-- name: GetPrice :one
SELECT id,title,benefit, price
FROM upgrade_price
WHERE is_choose = TRUE 
LIMIT 1; 

-- name: MakeOwner :one
INSERT INTO accounts(
    user_id,
    fullname,
    country,
    language,
    role_id,
    url_avatar,
    url_background_profile
) VALUES (
    $1, $2, $3, $4, 2, $5, $6
) RETURNING * ;

-- name: IsUpgradeAccount :one
SELECT accounts.* FROM accounts 
JOIN users ON accounts.user_id = users.id
WHERE users.username = $1 ORDER BY accounts.id LIMIT 1;