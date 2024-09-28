-- name: CreatePost :one
INSERT INTO posts(
    post_type_id,
    user_id,
    description,
    date_create_post
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1 AND post_type_id != 2 LIMIT 1;

-- name: ListPost :many
SELECT * FROM posts
WHERE post_type_id != 2
ORDER BY id DESC
LIMIT $1
OFFSET $2;

-- name: UpdatePost :one
UPDATE posts
SET description = $2
WHERE id = $1
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;

