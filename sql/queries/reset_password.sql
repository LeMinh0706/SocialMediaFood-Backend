-- name: GetUserByEmail :one
SELECT id, username FROM users
WHERE email = $1
LIMIT 1;

-- name: GetRequestByUUID :one
SELECT * FROM reset_password 
WHERE id = $1;

-- name: CreateRequestPassword :one
INSERT INTO reset_password (
    id, 
    user_id,
    expires_at
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: UpdatePassword :one
UPDATE users SET hash_password = $2
WHERE id = $1 
RETURNING id, created_at;

-- name: GetCheckAction :one
SELECT * FROM reset_password
WHERE user_id = $1
ORDER BY expires_at DESC
LIMIT 1;

-- name: GetUserById :one
SELECT id, username FROM users
WHERE id = $1
LIMIT 1;

-- name: UpdateActive :exec
UPDATE reset_password SET is_active = TRUE
WHERE id = $1;