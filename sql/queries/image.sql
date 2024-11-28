-- name: AddImagePost :one
INSERT INTO post_image (
    url_image,
    post_id
) VALUES (
    $1, $2
) RETURNING * ;

-- name: GetImage :one
SELECT * FROM post_image
WHERE id = $1;

-- name: GetImageComment :one
SELECT * FROM post_image
WHERE post_id = $1
ORDER BY id DESC
LIMIT 1;

-- name: GetImagePost :many
SELECT * FROM post_image 
WHERE post_id = $1;

-- name: DeleteImagePost :exec 
DELETE FROM post_image
WHERE id = $1;

-- name: DeleteImageComment :exec 
DELETE FROM post_image
WHERE post_id = $1;

-- name: UpdateImagePost :one
UPDATE post_image SET url_image = $2
WHERE post_id = $1
RETURNING *;

-- name: GetListImage :many
SELECT * FROM post_image
ORDER BY id DESC
LIMIT $1
OFFSET $2;