-- name: CreateFollow :one
INSERT INTO follower ( 
    from_follow,
    to_follow,
    status
) VALUES ($1, $2, $3)
RETURNING *;

-- name: GetFollowStatus :one
SELECT * FROM follower
WHERE from_follow = $1 AND to_follow = $2;

-- name: UpdateFriend :exec
UPDATE follower SET status = 'friend'
WHERE (from_follow = $1 AND to_follow = $2) OR (from_follow = $2 AND to_follow = $1);

-- name: DeleteFollow :exec
DELETE FROM follower
WHERE (from_follow = $1 AND to_follow = $2) OR (from_follow = $2 AND to_follow = $1);


