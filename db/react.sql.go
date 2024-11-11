// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: react.sql

package db

import (
	"context"
)

const countReactPost = `-- name: CountReactPost :one
SELECT count(id) FROM react_post
WHERE post_id = $1
`

func (q *Queries) CountReactPost(ctx context.Context, postID int64) (int64, error) {
	row := q.db.QueryRow(ctx, countReactPost, postID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createReact = `-- name: CreateReact :one
INSERT INTO react_post (
    account_id,
    post_id, 
    state
) VALUES (
    $1, $2, $3
) RETURNING id, account_id, post_id, state
`

type CreateReactParams struct {
	AccountID int64 `json:"account_id"`
	PostID    int64 `json:"post_id"`
	State     int32 `json:"state"`
}

func (q *Queries) CreateReact(ctx context.Context, arg CreateReactParams) (ReactPost, error) {
	row := q.db.QueryRow(ctx, createReact, arg.AccountID, arg.PostID, arg.State)
	var i ReactPost
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.PostID,
		&i.State,
	)
	return i, err
}

const deleteReact = `-- name: DeleteReact :exec
DELETE FROM react_post
WHERE post_id = $1 AND account_id = $2
`

type DeleteReactParams struct {
	PostID    int64 `json:"post_id"`
	AccountID int64 `json:"account_id"`
}

func (q *Queries) DeleteReact(ctx context.Context, arg DeleteReactParams) error {
	_, err := q.db.Exec(ctx, deleteReact, arg.PostID, arg.AccountID)
	return err
}

const getFavorite = `-- name: GetFavorite :many
SELECT post_id FROM react_post
WHERE account_id = $1
ORDER BY id DESC
LIMIT $2
OFFSET $3
`

type GetFavoriteParams struct {
	AccountID int64 `json:"account_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *Queries) GetFavorite(ctx context.Context, arg GetFavoriteParams) ([]int64, error) {
	rows, err := q.db.Query(ctx, getFavorite, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var post_id int64
		if err := rows.Scan(&post_id); err != nil {
			return nil, err
		}
		items = append(items, post_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getListReact = `-- name: GetListReact :many
SELECT id, account_id FROM react_post
WHERE post_id = $1
LIMIT $2
OFFSET $3
`

type GetListReactParams struct {
	PostID int64 `json:"post_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetListReactRow struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
}

func (q *Queries) GetListReact(ctx context.Context, arg GetListReactParams) ([]GetListReactRow, error) {
	rows, err := q.db.Query(ctx, getListReact, arg.PostID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetListReactRow{}
	for rows.Next() {
		var i GetListReactRow
		if err := rows.Scan(&i.ID, &i.AccountID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getReact = `-- name: GetReact :one
SELECT id, account_id, post_id, state FROM react_post
WHERE account_id = $1 AND post_id = $2
`

type GetReactParams struct {
	AccountID int64 `json:"account_id"`
	PostID    int64 `json:"post_id"`
}

func (q *Queries) GetReact(ctx context.Context, arg GetReactParams) (ReactPost, error) {
	row := q.db.QueryRow(ctx, getReact, arg.AccountID, arg.PostID)
	var i ReactPost
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.PostID,
		&i.State,
	)
	return i, err
}

const listAccountReact = `-- name: ListAccountReact :many
SELECT account_id FROM react_post
WHERE post_id = $1
`

func (q *Queries) ListAccountReact(ctx context.Context, postID int64) ([]int64, error) {
	rows, err := q.db.Query(ctx, listAccountReact, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var account_id int64
		if err := rows.Scan(&account_id); err != nil {
			return nil, err
		}
		items = append(items, account_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateState = `-- name: UpdateState :one
UPDATE react_post SET state = $3
WHERE post_id = $1 AND account_id = $2
RETURNING id, account_id, post_id, state
`

type UpdateStateParams struct {
	PostID    int64 `json:"post_id"`
	AccountID int64 `json:"account_id"`
	State     int32 `json:"state"`
}

func (q *Queries) UpdateState(ctx context.Context, arg UpdateStateParams) (ReactPost, error) {
	row := q.db.QueryRow(ctx, updateState, arg.PostID, arg.AccountID, arg.State)
	var i ReactPost
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.PostID,
		&i.State,
	)
	return i, err
}
