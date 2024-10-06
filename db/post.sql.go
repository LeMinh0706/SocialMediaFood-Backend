// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: post.sql

package db

import (
	"context"
	"database/sql"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts(
    post_type_id,
    user_id,
    description,
    date_create_post
) VALUES (
    $1, $2, $3, $4
) RETURNING id, post_type_id, user_id, post_top_id, description, date_create_post, is_banned, is_deleted
`

type CreatePostParams struct {
	PostTypeID     int32          `json:"post_type_id"`
	UserID         int64          `json:"user_id"`
	Description    sql.NullString `json:"description"`
	DateCreatePost int64          `json:"date_create_post"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.PostTypeID,
		arg.UserID,
		arg.Description,
		arg.DateCreatePost,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.PostTypeID,
		&i.UserID,
		&i.PostTopID,
		&i.Description,
		&i.DateCreatePost,
		&i.IsBanned,
		&i.IsDeleted,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1
`

func (q *Queries) DeletePost(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePost, id)
	return err
}

const getPost = `-- name: GetPost :one
SELECT id, post_type_id, user_id, post_top_id, description, date_create_post, is_banned, is_deleted FROM posts
WHERE id = $1 AND post_type_id != 2 AND (is_banned != true AND is_deleted != true) LIMIT 1
`

func (q *Queries) GetPost(ctx context.Context, id int64) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.PostTypeID,
		&i.UserID,
		&i.PostTopID,
		&i.Description,
		&i.DateCreatePost,
		&i.IsBanned,
		&i.IsDeleted,
	)
	return i, err
}

const listPost = `-- name: ListPost :many
SELECT id, post_type_id, user_id, post_top_id, description, date_create_post, is_banned, is_deleted FROM posts
WHERE post_type_id != 2 AND is_banned = false AND is_deleted = false
ORDER BY id DESC
LIMIT $1
OFFSET $2
`

type ListPostParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPost(ctx context.Context, arg ListPostParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPost, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.PostTypeID,
			&i.UserID,
			&i.PostTopID,
			&i.Description,
			&i.DateCreatePost,
			&i.IsBanned,
			&i.IsDeleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePost = `-- name: UpdatePost :one
UPDATE posts
SET description = $2
WHERE id = $1
RETURNING id, post_type_id, user_id, post_top_id, description, date_create_post, is_banned, is_deleted
`

type UpdatePostParams struct {
	ID          int64          `json:"id"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost, arg.ID, arg.Description)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.PostTypeID,
		&i.UserID,
		&i.PostTopID,
		&i.Description,
		&i.DateCreatePost,
		&i.IsBanned,
		&i.IsDeleted,
	)
	return i, err
}
