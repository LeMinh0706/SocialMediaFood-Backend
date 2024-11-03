-- name: CreateOwnerBranch :one
INSERT INTO locate (
    account_id,
    location
) VALUES (
    $1, ST_GeomFromText($2,4326)
) RETURNING id, account_id, ST_X(location::geometry) AS lng, ST_Y(location::geometry) AS lat;