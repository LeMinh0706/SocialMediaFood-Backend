-- name: CreateImagePost :one
INSERT INTO post_image(
    post_id,
    url_image
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetImagePost :many
SELECT * FROM post_image
WHERE post_id = $1;