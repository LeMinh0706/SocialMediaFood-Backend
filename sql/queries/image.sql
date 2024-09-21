-- name: CreateImagePost :one
INSERT INTO post_image(
    post_id,
    url_image
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetImagePost :one 
SELECT * FROM post_image
WHERE post_id = $1 AND url_image = $2
LIMIT 1;