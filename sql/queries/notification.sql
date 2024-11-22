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
WHERE id = $1;

-- name: GetListNoti :many 
SELECT * FROM (
    SELECT DISTINCT ON (n.post_id, n.type_id) id, user_action_id, created_at
    FROM notification n
    WHERE n.account_id = $1
    ORDER BY n.post_id, n.type_id, n.created_at DESC
) sub
ORDER BY sub.created_at DESC
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