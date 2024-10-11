-- name: CreateReact :one
INSERT INTO react_post(
    post_id,
    user_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: DeleteReact :exec
DELETE FROM react_post
WHERE id = $1;

-- name: GetReact :one
SELECT * FROM react_post
WHERE post_id = $1 AND user_id = $2;