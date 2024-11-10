-- name: Register :one
INSERT INTO users(
    username, 
    email,
    hash_password
) VALUES (
    $1, $2, $3
) RETURNING id, username, email, created_at;

-- name: Login :one
SELECT id, username, hash_password FROM users
WHERE username = $1;

-- name: UpdatePassword :one
UPDATE users SET hash_password = $2
WHERE id = $1 
RETURNING id, created_at;