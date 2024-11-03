// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: accounts.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAccounts = `-- name: CreateAccounts :one
INSERT INTO accounts(
    user_id,
    fullname,
    gender,
    country,
    language,
    role_id,
    url_avatar,
    url_background_profile,
    is_upgrade
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, false
) RETURNING id, user_id, fullname, url_avatar, url_background_profile, gender, country, language, address, is_deleted, role_id, is_upgrade, banned, introduce
`

type CreateAccountsParams struct {
	UserID               int64       `json:"user_id"`
	Fullname             string      `json:"fullname"`
	Gender               pgtype.Int4 `json:"gender"`
	Country              pgtype.Text `json:"country"`
	Language             pgtype.Text `json:"language"`
	RoleID               int32       `json:"role_id"`
	UrlAvatar            string      `json:"url_avatar"`
	UrlBackgroundProfile string      `json:"url_background_profile"`
}

func (q *Queries) CreateAccounts(ctx context.Context, arg CreateAccountsParams) (Account, error) {
	row := q.db.QueryRow(ctx, createAccounts,
		arg.UserID,
		arg.Fullname,
		arg.Gender,
		arg.Country,
		arg.Language,
		arg.RoleID,
		arg.UrlAvatar,
		arg.UrlBackgroundProfile,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Fullname,
		&i.UrlAvatar,
		&i.UrlBackgroundProfile,
		&i.Gender,
		&i.Country,
		&i.Language,
		&i.Address,
		&i.IsDeleted,
		&i.RoleID,
		&i.IsUpgrade,
		&i.Banned,
		&i.Introduce,
	)
	return i, err
}

const getAccountById = `-- name: GetAccountById :one
SELECT id, user_id, fullname, url_avatar, url_background_profile, gender, country, language, address, is_deleted, role_id, is_upgrade, banned, introduce FROM accounts
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetAccountById(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRow(ctx, getAccountById, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Fullname,
		&i.UrlAvatar,
		&i.UrlBackgroundProfile,
		&i.Gender,
		&i.Country,
		&i.Language,
		&i.Address,
		&i.IsDeleted,
		&i.RoleID,
		&i.IsUpgrade,
		&i.Banned,
		&i.Introduce,
	)
	return i, err
}

const getAccountByUserId = `-- name: GetAccountByUserId :many
SELECT id, user_id, fullname, url_avatar, url_background_profile, gender, country, language, role_id, address, is_upgrade
FROM accounts
WHERE user_id = $1
ORDER BY id
`

type GetAccountByUserIdRow struct {
	ID                   int64       `json:"id"`
	UserID               int64       `json:"user_id"`
	Fullname             string      `json:"fullname"`
	UrlAvatar            string      `json:"url_avatar"`
	UrlBackgroundProfile string      `json:"url_background_profile"`
	Gender               pgtype.Int4 `json:"gender"`
	Country              pgtype.Text `json:"country"`
	Language             pgtype.Text `json:"language"`
	RoleID               int32       `json:"role_id"`
	Address              pgtype.Text `json:"address"`
	IsUpgrade            pgtype.Bool `json:"is_upgrade"`
}

func (q *Queries) GetAccountByUserId(ctx context.Context, userID int64) ([]GetAccountByUserIdRow, error) {
	rows, err := q.db.Query(ctx, getAccountByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountByUserIdRow{}
	for rows.Next() {
		var i GetAccountByUserIdRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Fullname,
			&i.UrlAvatar,
			&i.UrlBackgroundProfile,
			&i.Gender,
			&i.Country,
			&i.Language,
			&i.RoleID,
			&i.Address,
			&i.IsUpgrade,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
