-- name: CreatePost :one
INSERT INTO posts (
    post_type_id,
    account_id,
    description,
    location
) VALUES (
    $1, $2, $3, ST_GeomFromText($4,4326)
) RETURNING id, post_type_id, account_id, description, ST_X(location::geometry) AS lng, ST_Y(location::geometry) AS lat, created_at;

-- name: UpdatePost :one
UPDATE posts SET description = $2
WHERE id = $1
RETURNING id, post_type_id, account_id, description, ST_X(location::geometry) AS lng, ST_Y(location::geometry) AS lat, created_at;

-- name: GetListPost :many
SELECT id
FROM posts
WHERE is_deleted != TRUE AND is_banned != TRUE AND post_type_id != 9
ORDER BY created_at DESC
LIMIT $1 
OFFSET $2;

-- name: GetHomePagePost :many
WITH posts_in_range AS (
    SELECT p.id, p.created_at, f.status
    FROM posts p
    LEFT JOIN follower as f ON p.account_id = f.to_follow AND f.from_follow = $1
    WHERE (f.from_follow = $1 OR f.from_follow IS NULL)
      AND is_deleted != TRUE
      AND is_banned != TRUE
      AND post_type_id != 9
    ORDER BY p.created_at DESC
    LIMIT $2 OFFSET $3
),
has_friend_posts AS (
    SELECT COUNT(*) AS friend_count
    FROM posts_in_range
    WHERE status IN ('friend','request')
)
SELECT p.id
FROM posts_in_range p, has_friend_posts h
ORDER BY 
    CASE
        WHEN h.friend_count > 0 AND p.status = 'friend' THEN 1
        WHEN h.friend_count > 0 AND p.status = 'request' THEN 2
        ELSE 3
    END,
    p.created_at DESC;

-- name: GetPost :one
SELECT id, post_type_id, account_id, description, ST_X(location::geometry) AS lng, ST_Y(location::geometry) AS lat, created_at
FROM posts 
WHERE id = $1 AND is_deleted != TRUE AND is_banned != TRUE AND post_type_id != 9;

-- name: DeletePost :exec
UPDATE posts SET is_deleted = TRUE
WHERE id = $1;

-- name: GetPersonPost :many
SELECT id FROM posts 
WHERE account_id = $1 AND is_deleted != TRUE AND is_banned != TRUE AND post_type_id != 9
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

--comment

-- name: GetListComment :many
SELECT id
FROM posts
WHERE post_top_id = $1 AND post_type_id = 9  
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetComment :one
SELECT id, post_type_id, account_id, post_top_id, description, created_at 
FROM posts
WHERE id = $1;

-- name: CreateComment :one
INSERT INTO posts (
    post_type_id,
    account_id,
    post_top_id,
    description
) VALUES (
    9, $1, $2, $3 
) RETURNING id, post_type_id, account_id, post_top_id, description, created_at;

-- name: UpdateComment :one
UPDATE posts SET description = $2
WHERE id = $1
RETURNING id, post_type_id, account_id, post_top_id, description, created_at;

-- name: DeleteComment :exec
DELETE FROM posts 
WHERE id = $1;

-- name: CountComment :one
SELECT count(id) FROM posts
WHERE post_top_id = $1;

-- name: GetPostInLocate :many
SELECT id, post_type_id, account_id, description, ST_X(location::geometry) AS lng, ST_Y(location::geometry) AS lat, created_at
FROM posts
WHERE is_banned != TRUE 
AND is_deleted != TRUE
AND ST_DWithin(location, ST_SetSRID(ST_MakePoint($1, $2), 4326), $3); 