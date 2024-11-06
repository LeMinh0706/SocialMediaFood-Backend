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
SELECT id, user_id, fullname, url_avatar, url_background_profile, role_id FROM accounts
WHERE id = $1
LIMIT 1;

-- name: UpdateName :one
UPDATE accounts SET fullname = $2
WHERE id = $1
RETURNING id, user_id, fullname;
