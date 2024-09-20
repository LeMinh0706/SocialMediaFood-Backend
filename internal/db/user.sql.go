// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(
    email,
    hash_pashword,
    fullname,
    gender,
    role_id,
    date_create_account
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id, email, hash_pashword, fullname, gender, country, language, url_avatar, role_id, url_background_profile, date_create_account
`

type CreateUserParams struct {
	Email             sql.NullString `json:"email"`
	HashPashword      string         `json:"hash_pashword"`
	Fullname          string         `json:"fullname"`
	Gender            int32          `json:"gender"`
	RoleID            int32          `json:"role_id"`
	DateCreateAccount int64          `json:"date_create_account"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Email,
		arg.HashPashword,
		arg.Fullname,
		arg.Gender,
		arg.RoleID,
		arg.DateCreateAccount,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashPashword,
		&i.Fullname,
		&i.Gender,
		&i.Country,
		&i.Language,
		&i.UrlAvatar,
		&i.RoleID,
		&i.UrlBackgroundProfile,
		&i.DateCreateAccount,
	)
	return i, err
}

const getListUser = `-- name: GetListUser :many
SELECT id, email, fullname, gender, role_id, date_create_account FROM users
ORDER BY id DESC
LIMIT $1 
OFFSET $2
`

type GetListUserParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetListUserRow struct {
	ID                int64          `json:"id"`
	Email             sql.NullString `json:"email"`
	Fullname          string         `json:"fullname"`
	Gender            int32          `json:"gender"`
	RoleID            int32          `json:"role_id"`
	DateCreateAccount int64          `json:"date_create_account"`
}

func (q *Queries) GetListUser(ctx context.Context, arg GetListUserParams) ([]GetListUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getListUser, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetListUserRow{}
	for rows.Next() {
		var i GetListUserRow
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Fullname,
			&i.Gender,
			&i.RoleID,
			&i.DateCreateAccount,
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

const getUser = `-- name: GetUser :one
SELECT id, email, fullname, gender, role_id, date_create_account FROM users 
WHERE id = $1
`

type GetUserRow struct {
	ID                int64          `json:"id"`
	Email             sql.NullString `json:"email"`
	Fullname          string         `json:"fullname"`
	Gender            int32          `json:"gender"`
	RoleID            int32          `json:"role_id"`
	DateCreateAccount int64          `json:"date_create_account"`
}

func (q *Queries) GetUser(ctx context.Context, id int64) (GetUserRow, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i GetUserRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Fullname,
		&i.Gender,
		&i.RoleID,
		&i.DateCreateAccount,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users 
SET fullname = $2,
gender = $3
WHERE id = $1
RETURNING id, email, fullname, gender, role_id, date_create_account
`

type UpdateUserParams struct {
	ID       int64  `json:"id"`
	Fullname string `json:"fullname"`
	Gender   int32  `json:"gender"`
}

type UpdateUserRow struct {
	ID                int64          `json:"id"`
	Email             sql.NullString `json:"email"`
	Fullname          string         `json:"fullname"`
	Gender            int32          `json:"gender"`
	RoleID            int32          `json:"role_id"`
	DateCreateAccount int64          `json:"date_create_account"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.ID, arg.Fullname, arg.Gender)
	var i UpdateUserRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Fullname,
		&i.Gender,
		&i.RoleID,
		&i.DateCreateAccount,
	)
	return i, err
}
