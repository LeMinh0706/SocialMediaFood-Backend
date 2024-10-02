-- name: CreateUser :one
INSERT INTO users(
    email,
    hash_pashword,
    username,
    fullname,
    gender,
    role_id,
    date_create_account
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username LIKE $1 LIMIT 1;

-- name: GetUserById :one
SELECT id, fullname, role_id FROM users 
WHERE id = $1 LIMIT 1;

-- name: GetListUser :many
SELECT id, email, fullname, gender, role_id, date_create_account FROM users
ORDER BY id DESC
LIMIT $1 
OFFSET $2;

-- name: UpdateUser :one
UPDATE users 
SET fullname = $2,
gender = $3
WHERE id = $1
RETURNING id, email, fullname, gender, role_id, date_create_account;

