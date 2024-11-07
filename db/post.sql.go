// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: post.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createComment = `-- name: CreateComment :one
INSERT INTO posts (
    post_type_id,
    account_id,
    post_top_id,
    description
) VALUES (
    9, $1, $2, $3 
) RETURNING id, account_id, post_top_id, description, created_at
`

type CreateCommentParams struct {
	AccountID   int64       `json:"account_id"`
	PostTopID   pgtype.Int8 `json:"post_top_id"`
	Description pgtype.Text `json:"description"`
}

type CreateCommentRow struct {
	ID          int64              `json:"id"`
	AccountID   int64              `json:"account_id"`
	PostTopID   pgtype.Int8        `json:"post_top_id"`
	Description pgtype.Text        `json:"description"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (CreateCommentRow, error) {
	row := q.db.QueryRow(ctx, createComment, arg.AccountID, arg.PostTopID, arg.Description)
	var i CreateCommentRow
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.PostTopID,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
    post_type_id,
    account_id,
    description,
    location
) VALUES (
    $1, $2, $3, ST_GeomFromText($4,4326)
) RETURNING id, post_type_id, account_id, description, ST_X(location::geometry) AS lng, ST_Y(location::geometry) AS lat, created_at
`

type CreatePostParams struct {
	PostTypeID     int32       `json:"post_type_id"`
	AccountID      int64       `json:"account_id"`
	Description    pgtype.Text `json:"description"`
	StGeomfromtext interface{} `json:"st_geomfromtext"`
}

type CreatePostRow struct {
	ID          int64              `json:"id"`
	PostTypeID  int32              `json:"post_type_id"`
	AccountID   int64              `json:"account_id"`
	Description pgtype.Text        `json:"description"`
	Lng         interface{}        `json:"lng"`
	Lat         interface{}        `json:"lat"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (CreatePostRow, error) {
	row := q.db.QueryRow(ctx, createPost,
		arg.PostTypeID,
		arg.AccountID,
		arg.Description,
		arg.StGeomfromtext,
	)
	var i CreatePostRow
	err := row.Scan(
		&i.ID,
		&i.PostTypeID,
		&i.AccountID,
		&i.Description,
		&i.Lng,
		&i.Lat,
		&i.CreatedAt,
	)
	return i, err
}

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM posts 
WHERE id = $1
`

func (q *Queries) DeleteComment(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteComment, id)
	return err
}

const deletePost = `-- name: DeletePost :exec
UPDATE posts SET is_deleted = TRUE
WHERE id = $1
`

func (q *Queries) DeletePost(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deletePost, id)
	return err
}

const getComment = `-- name: GetComment :one
SELECT id, account_id, post_top_id, description, created_at 
FROM posts
WHERE id = $1
`

type GetCommentRow struct {
	ID          int64              `json:"id"`
	AccountID   int64              `json:"account_id"`
	PostTopID   pgtype.Int8        `json:"post_top_id"`
	Description pgtype.Text        `json:"description"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) GetComment(ctx context.Context, id int64) (GetCommentRow, error) {
	row := q.db.QueryRow(ctx, getComment, id)
	var i GetCommentRow
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.PostTopID,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const getListComment = `-- name: GetListComment :many

SELECT id
FROM posts
WHERE post_top_id = $1 AND post_type_id = 9  
ORDER BY created_at DESC
LIMIT $2
OFFSET $3
`

type GetListCommentParams struct {
	PostTopID pgtype.Int8 `json:"post_top_id"`
	Limit     int32       `json:"limit"`
	Offset    int32       `json:"offset"`
}

// comment
func (q *Queries) GetListComment(ctx context.Context, arg GetListCommentParams) ([]int64, error) {
	rows, err := q.db.Query(ctx, getListComment, arg.PostTopID, arg.Limit, arg.Offset)
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

const getListPost = `-- name: GetListPost :many
SELECT id
FROM posts
WHERE is_deleted != TRUE AND is_banned != TRUE AND post_type_id != 9
ORDER BY created_at DESC
LIMIT $1 
OFFSET $2
`

type GetListPostParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetListPost(ctx context.Context, arg GetListPostParams) ([]int64, error) {
	rows, err := q.db.Query(ctx, getListPost, arg.Limit, arg.Offset)
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

const getPersonPost = `-- name: GetPersonPost :many
SELECT id FROM posts 
WHERE account_id = $1 AND is_deleted != TRUE AND is_banned != TRUE AND post_type_id != 9
ORDER BY created_at DESC
LIMIT $2
OFFSET $3
`

type GetPersonPostParams struct {
	AccountID int64 `json:"account_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *Queries) GetPersonPost(ctx context.Context, arg GetPersonPostParams) ([]int64, error) {
	rows, err := q.db.Query(ctx, getPersonPost, arg.AccountID, arg.Limit, arg.Offset)
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

const getPost = `-- name: GetPost :one
SELECT id, post_type_id, account_id, description, ST_X(location::geometry) AS lng, ST_Y(location::geometry) AS lat, created_at
FROM posts 
WHERE id = $1 AND is_deleted != TRUE AND is_banned != TRUE AND post_type_id != 9
`

type GetPostRow struct {
	ID          int64              `json:"id"`
	PostTypeID  int32              `json:"post_type_id"`
	AccountID   int64              `json:"account_id"`
	Description pgtype.Text        `json:"description"`
	Lng         interface{}        `json:"lng"`
	Lat         interface{}        `json:"lat"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) GetPost(ctx context.Context, id int64) (GetPostRow, error) {
	row := q.db.QueryRow(ctx, getPost, id)
	var i GetPostRow
	err := row.Scan(
		&i.ID,
		&i.PostTypeID,
		&i.AccountID,
		&i.Description,
		&i.Lng,
		&i.Lat,
		&i.CreatedAt,
	)
	return i, err
}

const updateComment = `-- name: UpdateComment :one
UPDATE posts SET description = $2
WHERE id = $1
RETURNING id, account_id, post_top_id, description, created_at
`

type UpdateCommentParams struct {
	ID          int64       `json:"id"`
	Description pgtype.Text `json:"description"`
}

type UpdateCommentRow struct {
	ID          int64              `json:"id"`
	AccountID   int64              `json:"account_id"`
	PostTopID   pgtype.Int8        `json:"post_top_id"`
	Description pgtype.Text        `json:"description"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) UpdateComment(ctx context.Context, arg UpdateCommentParams) (UpdateCommentRow, error) {
	row := q.db.QueryRow(ctx, updateComment, arg.ID, arg.Description)
	var i UpdateCommentRow
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.PostTopID,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const updatePost = `-- name: UpdatePost :one
UPDATE posts SET description = $2
WHERE id = $1
RETURNING id, post_type_id, account_id, description, ST_X(location::geometry) AS lng, ST_Y(location::geometry) AS lat, created_at
`

type UpdatePostParams struct {
	ID          int64       `json:"id"`
	Description pgtype.Text `json:"description"`
}

type UpdatePostRow struct {
	ID          int64              `json:"id"`
	PostTypeID  int32              `json:"post_type_id"`
	AccountID   int64              `json:"account_id"`
	Description pgtype.Text        `json:"description"`
	Lng         interface{}        `json:"lng"`
	Lat         interface{}        `json:"lat"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (UpdatePostRow, error) {
	row := q.db.QueryRow(ctx, updatePost, arg.ID, arg.Description)
	var i UpdatePostRow
	err := row.Scan(
		&i.ID,
		&i.PostTypeID,
		&i.AccountID,
		&i.Description,
		&i.Lng,
		&i.Lat,
		&i.CreatedAt,
	)
	return i, err
}
