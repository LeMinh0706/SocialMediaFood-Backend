// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: notification.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createActionNoti = `-- name: CreateActionNoti :one
INSERT INTO notification (
    message,
    account_id,
    type_id,
    user_action_id
) VALUES (
    $1, $2, $3, $4
) RETURNING id, message, account_id, type_id, post_id, user_action_id, invoice_id, is_seen, created_at
`

type CreateActionNotiParams struct {
	Message      string `json:"message"`
	AccountID    int64  `json:"account_id"`
	TypeID       int32  `json:"type_id"`
	UserActionID int64  `json:"user_action_id"`
}

func (q *Queries) CreateActionNoti(ctx context.Context, arg CreateActionNotiParams) (Notification, error) {
	row := q.db.QueryRow(ctx, createActionNoti,
		arg.Message,
		arg.AccountID,
		arg.TypeID,
		arg.UserActionID,
	)
	var i Notification
	err := row.Scan(
		&i.ID,
		&i.Message,
		&i.AccountID,
		&i.TypeID,
		&i.PostID,
		&i.UserActionID,
		&i.InvoiceID,
		&i.IsSeen,
		&i.CreatedAt,
	)
	return i, err
}

const createPostNoti = `-- name: CreatePostNoti :one
INSERT INTO notification (
    message,
    account_id,
    type_id,
    post_id,
    user_action_id
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, message, account_id, type_id, post_id, user_action_id, invoice_id, is_seen, created_at
`

type CreatePostNotiParams struct {
	Message      string      `json:"message"`
	AccountID    int64       `json:"account_id"`
	TypeID       int32       `json:"type_id"`
	PostID       pgtype.Int8 `json:"post_id"`
	UserActionID int64       `json:"user_action_id"`
}

func (q *Queries) CreatePostNoti(ctx context.Context, arg CreatePostNotiParams) (Notification, error) {
	row := q.db.QueryRow(ctx, createPostNoti,
		arg.Message,
		arg.AccountID,
		arg.TypeID,
		arg.PostID,
		arg.UserActionID,
	)
	var i Notification
	err := row.Scan(
		&i.ID,
		&i.Message,
		&i.AccountID,
		&i.TypeID,
		&i.PostID,
		&i.UserActionID,
		&i.InvoiceID,
		&i.IsSeen,
		&i.CreatedAt,
	)
	return i, err
}

const deleteNoti = `-- name: DeleteNoti :exec
DELETE FROM notification 
WHERE id = $1
`

func (q *Queries) DeleteNoti(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteNoti, id)
	return err
}

const getListNoti = `-- name: GetListNoti :many
SELECT id FROM notification
WHERE account_id = $1
ORDER BY id DESC
LIMIT $2
OFFSET $3
`

type GetListNotiParams struct {
	AccountID int64 `json:"account_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *Queries) GetListNoti(ctx context.Context, arg GetListNotiParams) ([]int64, error) {
	rows, err := q.db.Query(ctx, getListNoti, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNotification = `-- name: GetNotification :one
SELECT id, message, account_id, type_id, post_id, user_action_id, invoice_id, is_seen, created_at FROM notification
WHERE account_id = $1
LIMIT 1
`

func (q *Queries) GetNotification(ctx context.Context, accountID int64) (Notification, error) {
	row := q.db.QueryRow(ctx, getNotification, accountID)
	var i Notification
	err := row.Scan(
		&i.ID,
		&i.Message,
		&i.AccountID,
		&i.TypeID,
		&i.PostID,
		&i.UserActionID,
		&i.InvoiceID,
		&i.IsSeen,
		&i.CreatedAt,
	)
	return i, err
}

const updateSeen = `-- name: UpdateSeen :exec
UPDATE notification SET is_seen = true
WHERE id = $1
`

func (q *Queries) UpdateSeen(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, updateSeen, id)
	return err
}

const updateSeenAll = `-- name: UpdateSeenAll :exec
UPDATE notification SET is_seen = true
WHERE account_id = $1
`

func (q *Queries) UpdateSeenAll(ctx context.Context, accountID int64) error {
	_, err := q.db.Exec(ctx, updateSeenAll, accountID)
	return err
}
