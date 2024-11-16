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
SELECT id, user_id, fullname, url_avatar, url_background_profile, role_id FROM accounts
WHERE id = $1
LIMIT 1
`

type GetAccountByIdRow struct {
	ID                   int64  `json:"id"`
	UserID               int64  `json:"user_id"`
	Fullname             string `json:"fullname"`
	UrlAvatar            string `json:"url_avatar"`
	UrlBackgroundProfile string `json:"url_background_profile"`
	RoleID               int32  `json:"role_id"`
}

func (q *Queries) GetAccountById(ctx context.Context, id int64) (GetAccountByIdRow, error) {
	row := q.db.QueryRow(ctx, getAccountById, id)
	var i GetAccountByIdRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Fullname,
		&i.UrlAvatar,
		&i.UrlBackgroundProfile,
		&i.RoleID,
	)
	return i, err
}

const getAccountByUserId = `-- name: GetAccountByUserId :many
SELECT id
FROM accounts
WHERE user_id = $1
ORDER BY id
`

func (q *Queries) GetAccountByUserId(ctx context.Context, userID int64) ([]int64, error) {
	rows, err := q.db.Query(ctx, getAccountByUserId, userID)
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

const getDetailAccount = `-- name: GetDetailAccount :one
SELECT id, user_id, fullname, url_avatar, url_background_profile, gender, country, language, address, is_deleted, role_id, is_upgrade, banned, introduce FROM accounts
WHERE id = $1
`

func (q *Queries) GetDetailAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRow(ctx, getDetailAccount, id)
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

const searchingAccounts = `-- name: SearchingAccounts :many
SELECT id, user_id, fullname, url_avatar, url_background_profile, role_id FROM accounts
WHERE fullname LIKE '%' || $1 || '%'
LIMIT $2
OFFSET $3
`

type SearchingAccountsParams struct {
	Column1 pgtype.Text `json:"column_1"`
	Limit   int32       `json:"limit"`
	Offset  int32       `json:"offset"`
}

type SearchingAccountsRow struct {
	ID                   int64  `json:"id"`
	UserID               int64  `json:"user_id"`
	Fullname             string `json:"fullname"`
	UrlAvatar            string `json:"url_avatar"`
	UrlBackgroundProfile string `json:"url_background_profile"`
	RoleID               int32  `json:"role_id"`
}

func (q *Queries) SearchingAccounts(ctx context.Context, arg SearchingAccountsParams) ([]SearchingAccountsRow, error) {
	rows, err := q.db.Query(ctx, searchingAccounts, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SearchingAccountsRow{}
	for rows.Next() {
		var i SearchingAccountsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Fullname,
			&i.UrlAvatar,
			&i.UrlBackgroundProfile,
			&i.RoleID,
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

const updateAvatar = `-- name: UpdateAvatar :one
UPDATE accounts SET url_avatar = $2
WHERE id = $1
RETURNING id, user_id, fullname, url_avatar, url_background_profile, role_id
`

type UpdateAvatarParams struct {
	ID        int64  `json:"id"`
	UrlAvatar string `json:"url_avatar"`
}

type UpdateAvatarRow struct {
	ID                   int64  `json:"id"`
	UserID               int64  `json:"user_id"`
	Fullname             string `json:"fullname"`
	UrlAvatar            string `json:"url_avatar"`
	UrlBackgroundProfile string `json:"url_background_profile"`
	RoleID               int32  `json:"role_id"`
}

func (q *Queries) UpdateAvatar(ctx context.Context, arg UpdateAvatarParams) (UpdateAvatarRow, error) {
	row := q.db.QueryRow(ctx, updateAvatar, arg.ID, arg.UrlAvatar)
	var i UpdateAvatarRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Fullname,
		&i.UrlAvatar,
		&i.UrlBackgroundProfile,
		&i.RoleID,
	)
	return i, err
}

const updateBackground = `-- name: UpdateBackground :one
UPDATE accounts SET url_background_profile = $2
WHERE id = $1
RETURNING id, user_id, fullname, url_avatar, url_background_profile, role_id
`

type UpdateBackgroundParams struct {
	ID                   int64  `json:"id"`
	UrlBackgroundProfile string `json:"url_background_profile"`
}

type UpdateBackgroundRow struct {
	ID                   int64  `json:"id"`
	UserID               int64  `json:"user_id"`
	Fullname             string `json:"fullname"`
	UrlAvatar            string `json:"url_avatar"`
	UrlBackgroundProfile string `json:"url_background_profile"`
	RoleID               int32  `json:"role_id"`
}

func (q *Queries) UpdateBackground(ctx context.Context, arg UpdateBackgroundParams) (UpdateBackgroundRow, error) {
	row := q.db.QueryRow(ctx, updateBackground, arg.ID, arg.UrlBackgroundProfile)
	var i UpdateBackgroundRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Fullname,
		&i.UrlAvatar,
		&i.UrlBackgroundProfile,
		&i.RoleID,
	)
	return i, err
}

const updateName = `-- name: UpdateName :one
UPDATE accounts SET fullname = $2
WHERE id = $1
RETURNING id, user_id, fullname
`

type UpdateNameParams struct {
	ID       int64  `json:"id"`
	Fullname string `json:"fullname"`
}

type UpdateNameRow struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"user_id"`
	Fullname string `json:"fullname"`
}

func (q *Queries) UpdateName(ctx context.Context, arg UpdateNameParams) (UpdateNameRow, error) {
	row := q.db.QueryRow(ctx, updateName, arg.ID, arg.Fullname)
	var i UpdateNameRow
	err := row.Scan(&i.ID, &i.UserID, &i.Fullname)
	return i, err
}

const upgradeSuccess = `-- name: UpgradeSuccess :one
UPDATE accounts SET is_upgrade = TRUE
WHERE id = $1
RETURNING id
`

func (q *Queries) UpgradeSuccess(ctx context.Context, id int64) (int64, error) {
	row := q.db.QueryRow(ctx, upgradeSuccess, id)
	err := row.Scan(&id)
	return id, err
}
