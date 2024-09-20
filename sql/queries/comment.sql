-- name: CreateComment :one
INSERT INTO posts(
    post_type_id,
    user_id,
    post_top_id,
    description,
    date_create_post
) VALUES (
    2, $1, $2, $3, $4
) RETURNING *;

-- name: GetComment :one
SELECT * FROM posts 
WHERE post_type_id = 2 AND
post_top_id = $1
LIMIT 1;

-- name: ListComment :many
SELECT * FROM posts 
WHERE post_type_id = 2 AND 
post_top_id = $1
LIMIT $2
OFFSET $3;