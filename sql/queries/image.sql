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

-- name: GetImagePost :many
SELECT * FROM post_image 
WHERE post_id = $1;

-- name: DeleteImagePost :exec 
DELETE FROM post_image
WHERE id = $1;