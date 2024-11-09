-- name: CreatePost :one
INSERT INTO posts (
    post_type_id,
    account_id,
    description,
    location
) VALUES (
    $1, $2, $3, ST_GeomFromText($4,4326)
) RETURNING id, post_type_id, account_id, description, ST_X(location::geometry) AS lng, ST_Y(location::geometry) AS lat, created_at;

-- name: UpdatePost :one
UPDATE posts SET description = $2
WHERE id = $1
RETURNING id, post_type_id, account_id, description, ST_X(location::geometry) AS lng, ST_Y(location::geometry) AS lat, created_at;

-- name: GetListPost :many
SELECT id
FROM posts
WHERE is_deleted != TRUE AND is_banned != TRUE AND post_type_id != 9
ORDER BY created_at DESC
LIMIT $1 
OFFSET $2;

-- name: GetPost :one
SELECT id, post_type_id, account_id, description, ST_X(location::geometry) AS lng, ST_Y(location::geometry) AS lat, created_at
FROM posts 
WHERE id = $1 AND is_deleted != TRUE AND is_banned != TRUE AND post_type_id != 9;

-- name: DeletePost :exec
UPDATE posts SET is_deleted = TRUE
WHERE id = $1;

-- name: GetPersonPost :many
SELECT id FROM posts 
WHERE account_id = $1 AND is_deleted != TRUE AND is_banned != TRUE AND post_type_id != 9
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

--comment

-- name: GetListComment :many
SELECT id
FROM posts
WHERE post_top_id = $1 AND post_type_id = 9  
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetComment :one
SELECT id, account_id, post_top_id, description, created_at 
FROM posts
WHERE id = $1;

-- name: CreateComment :one
INSERT INTO posts (
    post_type_id,
    account_id,
    post_top_id,
    description
) VALUES (
    9, $1, $2, $3 
) RETURNING id, account_id, post_top_id, description, created_at;

-- name: UpdateComment :one
UPDATE posts SET description = $2
WHERE id = $1
RETURNING id, account_id, post_top_id, description, created_at;

-- name: DeleteComment :exec
DELETE FROM posts 
WHERE id = $1;

-- name: CountComment :one
SELECT count(id) FROM posts
WHERE post_top_id = $1;