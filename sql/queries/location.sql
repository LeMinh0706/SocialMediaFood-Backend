-- name: CreateOwnerBranch :one
INSERT INTO locate (
    account_id,
    location
) VALUES (
    $1, ST_GeomFromText($2,4326)
) RETURNING id, account_id, ST_X(location::geometry) AS lng, ST_Y(location::geometry) AS lat;

-- name: GetLocation :many
SELECT id, ST_X(location::geometry) AS lng, ST_Y(location::geometry) AS lat
FROM locate
WHERE account_id = $1;

DELETE FROM locate
WHERE id = $1;