-- name: CreatePostNoti :one
INSERT INTO notification (
    message,
    account_id,
    type_id,
    post_id,
    user_action_id
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: CreateActionNoti :one
INSERT INTO notification (
    message,
    account_id,
    type_id,
    user_action_id
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetNotification :one
SELECT * FROM notification
WHERE account_id = $1
LIMIT 1;

-- name: GetListNoti :many 
SELECT id FROM notification
WHERE account_id = $1
ORDER BY id DESC
LIMIT $2
OFFSET $3;

-- name: UpdateSeen :exec
UPDATE notification SET is_seen = true
WHERE id = $1;

-- name: UpdateSeenAll :exec
UPDATE notification SET is_seen = true
WHERE account_id = $1;

-- name: DeleteNoti :exec
DELETE FROM notification 
WHERE id = $1;