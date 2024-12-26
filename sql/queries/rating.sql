-- name: CreateRating :one
INSERT INTO rating (
    from_account_id,
    to_account_id,
    star,
    content
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: DeleteRating :exec
DELETE FROM rating WHERE from_account_id = $1 AND to_account_id = $2;

-- name: UpdateRating :exec
UPDATE rating SET content = $3, star = $4
WHERE from_account_id = $1 AND to_account_id = $2;

-- name: GetListRating :many
SELECT r.*, a.fullname, a.url_avatar, a.url_background_profile FROM rating r JOIN accounts a
ON r.from_account_id = a.id
WHERE to_account_id = $1
LIMIT $2
OFFSET $3;
