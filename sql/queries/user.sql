-- name: CreateUser :one
INSERT INTO users(
    email,
    hash_pashword,
    fullname,
    gender,
    role_id,
    date_create_account
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;