-- name: CreateReact :one
INSERT INTO react_post (
    account_id,
    post_id, 
    state
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetReactPost :many
SELECT * FROM react_post
WHERE post_id = $1;

-- name: GetFavorite :many
SELECT post_id FROM react_post
WHERE account_id = $1
ORDER BY id DESC
LIMIT $2
OFFSET $3;