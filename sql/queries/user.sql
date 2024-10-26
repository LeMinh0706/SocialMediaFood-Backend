-- name: Register :one
INSERT INTO users(
    username, 
    email,
    hash_password
) VALUES (
    $1, $2, $3
) RETURNING id, username, email, created_at;

-- name: Login :one
SELECT id, username, hash_password, email, created_at FROM users
WHERE username = $1;