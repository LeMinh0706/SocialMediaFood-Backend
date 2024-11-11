-- name: CreateReact :one
INSERT INTO react_post (
    account_id,
    post_id, 
    state
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetReact :one
SELECT * FROM react_post
WHERE account_id = $1 AND post_id = $2;

-- name: GetListReact :many
SELECT id, account_id FROM react_post
WHERE post_id = $1
LIMIT $2
OFFSET $3;

-- name: CountReactPost :one
SELECT count(id) FROM react_post
WHERE post_id = $1;

-- name: GetFavorite :many
SELECT post_id FROM react_post
WHERE account_id = $1
ORDER BY id DESC
LIMIT $2
OFFSET $3;

-- name: UpdateState :one
UPDATE react_post SET state = $3
WHERE post_id = $1 AND account_id = $2
RETURNING * ;

-- name: DeleteReact :exec
DELETE FROM react_post
WHERE post_id = $1 AND account_id = $2;

-- name: ListAccountReact :many
SELECT account_id FROM react_post
WHERE post_id = $1;