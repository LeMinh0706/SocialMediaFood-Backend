-- name: CreatePost :one
INSERT INTO posts(
    post_type_id,
    user_id,
    post_top_id,
    description,
    date_create_post
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;