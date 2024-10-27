-- name: CreatePost :one
INSERT INTO posts (
    post_type_id,
    account_id,
    description,
    location
) VALUES (
    $1, $2, $3, ST_GeomFromText($4,4326)
) RETURNING id, post_type_id, account_id, description, ST_X(location::geometry) AS lng, ST_Y(location::geometry) AS lat, created_at;

-- name: CreateComment :one
INSERT INTO posts (
    post_type_id,
    account_id,
    post_top_id,
    description
) VALUES (
    9, $1, $2, $3 
) RETURNING * ;