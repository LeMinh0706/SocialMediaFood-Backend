-- name: CreateAccounts :one
INSERT INTO accounts(
    user_id,
    fullname,
    gender,
    country,
    language,
    type,
    url_avatar,
    url_background_profile
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING * ;

-- name: GetAccountByUserId :many
SELECT * FROM accounts
WHERE user_id = $1;

-- name: GetAccountById :one
SELECT * FROM accounts
WHERE id = $1
LIMIT 1;