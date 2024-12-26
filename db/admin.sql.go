// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: admin.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addUpgradePrice = `-- name: AddUpgradePrice :one
INSERT INTO upgrade_price (
    title,
    price,
    benefit
)VALUES ($1, $2, $3)
RETURNING id, title, benefit, price, is_choose, created_at
`

type AddUpgradePriceParams struct {
	Title   string         `json:"title"`
	Price   pgtype.Numeric `json:"price"`
	Benefit string         `json:"benefit"`
}

func (q *Queries) AddUpgradePrice(ctx context.Context, arg AddUpgradePriceParams) (UpgradePrice, error) {
	row := q.db.QueryRow(ctx, addUpgradePrice, arg.Title, arg.Price, arg.Benefit)
	var i UpgradePrice
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Benefit,
		&i.Price,
		&i.IsChoose,
		&i.CreatedAt,
	)
	return i, err
}

const banPost = `-- name: BanPost :exec
UPDATE posts SET is_banned = TRUE 
WHERE id = $1
`

func (q *Queries) BanPost(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, banPost, id)
	return err
}

const deleteReportPost = `-- name: DeleteReportPost :exec
DELETE FROM report_post WHERE post_id = $1
`

func (q *Queries) DeleteReportPost(ctx context.Context, postID int64) error {
	_, err := q.db.Exec(ctx, deleteReportPost, postID)
	return err
}

const getChoosePrice = `-- name: GetChoosePrice :one
SELECT id, title,benefit,price,created_at FROM upgrade_price
WHERE is_choose = TRUE
LIMIT 1
`

type GetChoosePriceRow struct {
	ID        int64              `json:"id"`
	Title     string             `json:"title"`
	Benefit   string             `json:"benefit"`
	Price     pgtype.Numeric     `json:"price"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) GetChoosePrice(ctx context.Context) (GetChoosePriceRow, error) {
	row := q.db.QueryRow(ctx, getChoosePrice)
	var i GetChoosePriceRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Benefit,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const getListPostReport = `-- name: GetListPostReport :many
SELECT post_id, COUNT(account_id), MAX(created_at)
FROM report_post
GROUP BY post_id
HAVING COUNT(account_id) >= 5
ORDER BY MAX(created_at) DESC
LIMIT $1
OFFSET $2
`

type GetListPostReportParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetListPostReportRow struct {
	PostID int64       `json:"post_id"`
	Count  int64       `json:"count"`
	Max    interface{} `json:"max"`
}

func (q *Queries) GetListPostReport(ctx context.Context, arg GetListPostReportParams) ([]GetListPostReportRow, error) {
	rows, err := q.db.Query(ctx, getListPostReport, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetListPostReportRow{}
	for rows.Next() {
		var i GetListPostReportRow
		if err := rows.Scan(&i.PostID, &i.Count, &i.Max); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getListUpgradePrice = `-- name: GetListUpgradePrice :many
SELECT id, title, benefit, price, is_choose, created_at FROM upgrade_price
ORDER BY id DESC
LIMIT $1
OFFSET $2
`

type GetListUpgradePriceParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetListUpgradePrice(ctx context.Context, arg GetListUpgradePriceParams) ([]UpgradePrice, error) {
	rows, err := q.db.Query(ctx, getListUpgradePrice, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UpgradePrice{}
	for rows.Next() {
		var i UpgradePrice
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Benefit,
			&i.Price,
			&i.IsChoose,
			&i.CreatedAt,
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

const getStatusQueue = `-- name: GetStatusQueue :one
SELECT status FROM upgrade_queue
WHERE account_id = $1
LIMIT 1
`

func (q *Queries) GetStatusQueue(ctx context.Context, accountID int64) (string, error) {
	row := q.db.QueryRow(ctx, getStatusQueue, accountID)
	var status string
	err := row.Scan(&status)
	return status, err
}

const getUpgradeQueue = `-- name: GetUpgradeQueue :many
SELECT account_id FROM upgrade_queue
WHERE status = 'pending'
ORDER BY created_at
LIMIT $1
OFFSET $2
`

type GetUpgradeQueueParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetUpgradeQueue(ctx context.Context, arg GetUpgradeQueueParams) ([]int64, error) {
	rows, err := q.db.Query(ctx, getUpgradeQueue, arg.Limit, arg.Offset)
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

const getUpgradeSuccess = `-- name: GetUpgradeSuccess :many
SELECT account_id FROM upgrade_queue
WHERE status = 'paid'
ORDER BY created_at
LIMIT $1
OFFSET $2
`

type GetUpgradeSuccessParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetUpgradeSuccess(ctx context.Context, arg GetUpgradeSuccessParams) ([]int64, error) {
	rows, err := q.db.Query(ctx, getUpgradeSuccess, arg.Limit, arg.Offset)
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

const isAdmin = `-- name: IsAdmin :one
SELECT role_id FROM accounts
WHERE user_id = $1
ORDER BY id ASC
LIMIT 1
`

func (q *Queries) IsAdmin(ctx context.Context, userID int64) (int32, error) {
	row := q.db.QueryRow(ctx, isAdmin, userID)
	var role_id int32
	err := row.Scan(&role_id)
	return role_id, err
}

const priceChoosing = `-- name: PriceChoosing :exec
UPDATE upgrade_price SET is_choose = TRUE 
WHERE id = $1
`

func (q *Queries) PriceChoosing(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, priceChoosing, id)
	return err
}

const reportPostDetail = `-- name: ReportPostDetail :many
SELECT r.account_id, r.issue_id, r.created_at, i.id, i.name, i.is_deleted FROM report_post r
LEFT JOIN issue_post i
ON r.issue_id = i.id 
WHERE r.post_id = $1  
ORDER BY r.created_at DESC 
LIMIT $2
OFFSET $3
`

type ReportPostDetailParams struct {
	PostID int64 `json:"post_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ReportPostDetailRow struct {
	AccountID int64              `json:"account_id"`
	IssueID   int32              `json:"issue_id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	ID        pgtype.Int4        `json:"id"`
	Name      pgtype.Text        `json:"name"`
	IsDeleted pgtype.Bool        `json:"is_deleted"`
}

func (q *Queries) ReportPostDetail(ctx context.Context, arg ReportPostDetailParams) ([]ReportPostDetailRow, error) {
	rows, err := q.db.Query(ctx, reportPostDetail, arg.PostID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ReportPostDetailRow{}
	for rows.Next() {
		var i ReportPostDetailRow
		if err := rows.Scan(
			&i.AccountID,
			&i.IssueID,
			&i.CreatedAt,
			&i.ID,
			&i.Name,
			&i.IsDeleted,
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

const unableChoose = `-- name: UnableChoose :exec
UPDATE upgrade_price SET is_choose = FALSE
`

func (q *Queries) UnableChoose(ctx context.Context) error {
	_, err := q.db.Exec(ctx, unableChoose)
	return err
}

const upgradeOwner = `-- name: UpgradeOwner :exec
UPDATE accounts SET is_upgrade = TRUE, role_id = 2
WHERE id = $1
`

func (q *Queries) UpgradeOwner(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, upgradeOwner, id)
	return err
}

const upgradeReject = `-- name: UpgradeReject :exec
DELETE FROM upgrade_queue WHERE account_id = $1
`

func (q *Queries) UpgradeReject(ctx context.Context, accountID int64) error {
	_, err := q.db.Exec(ctx, upgradeReject, accountID)
	return err
}

const upgradeStateQueue = `-- name: UpgradeStateQueue :exec
UPDATE upgrade_queue SET status = 'paid'
WHERE account_id = $1
`

func (q *Queries) UpgradeStateQueue(ctx context.Context, accountID int64) error {
	_, err := q.db.Exec(ctx, upgradeStateQueue, accountID)
	return err
}
