-- name: CreateFollow :one
INSERT INTO follower ( 
    from_follow,
    to_follow,
    status
) VALUES ($1, $2, $3)
RETURNING *;

-- name: GetFollowStatus :one
SELECT from_follow, to_follow, status FROM follower
WHERE from_follow = $1 AND to_follow = $2;

-- name: CountRequest :one
SELECT count(id) FROM follower 
WHERE from_follow = $1 AND status = 'request';

-- name: GetYourRequest :many
SELECT to_follow FROM follower
WHERE from_follow = $1 AND status = 'request'
ORDER BY id DESC
LIMIT $2
OFFSET $3;

-- name: CountFollower :one
SELECT count(id) FROM follower 
WHERE from_follow = $1 AND status = 'accept';

-- name: GetYourFollower :many
SELECT to_follow FROM follower
WHERE from_follow = $1 AND status = 'accept'
ORDER BY id DESC
LIMIT $2
OFFSET $3;

-- name: CountFriend :one
SELECT count(id) FROM follower 
WHERE from_follow = $1 AND status = 'friend';

-- name: GetYourFriend :many
SELECT to_follow FROM follower
WHERE from_follow = $1 AND status = 'friend'
ORDER BY id DESC
LIMIT $2
OFFSET $3;

-- name: UpdateFriend :exec
UPDATE follower SET status = 'friend'
WHERE (from_follow = $1 AND to_follow = $2) OR (from_follow = $2 AND to_follow = $1);

-- name: DeleteFollow :exec
DELETE FROM follower
WHERE (from_follow = $1 AND to_follow = $2) OR (from_follow = $2 AND to_follow = $1);



