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
SELECT id, user_id, fullname, url_avatar, url_background_profile, gender, country, language, role_id, address, is_upgrade
FROM accounts
WHERE user_id = $1
ORDER BY id;

-- name: GetAccountById :one
SELECT * FROM accounts
WHERE id = $1
LIMIT 1;
