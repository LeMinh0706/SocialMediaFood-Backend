-- name: GetListIssue :many
SELECT * FROM issue_post;

-- name: GetIssue :one
SELECT * FROM issue_post
WHERE id = $1
LIMIT 1;

-- name: CreateReport :one
INSERT INTO report_post (
    account_id,
    post_id,
    issue_id
) VALUES (
    $1, $2, $3
)RETURNING *;

-- name: GetYourReport :many
SELECT id, issue_id FROM report_post
WHERE account_id = $1 AND post_id = $2;