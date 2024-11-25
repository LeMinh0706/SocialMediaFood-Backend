// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID                   int64       `json:"id"`
	UserID               int64       `json:"user_id"`
	Fullname             string      `json:"fullname"`
	UrlAvatar            string      `json:"url_avatar"`
	UrlBackgroundProfile string      `json:"url_background_profile"`
	Gender               pgtype.Int4 `json:"gender"`
	Country              pgtype.Text `json:"country"`
	Language             pgtype.Text `json:"language"`
	Address              pgtype.Text `json:"address"`
	IsDeleted            bool        `json:"is_deleted"`
	RoleID               int32       `json:"role_id"`
	IsUpgrade            pgtype.Bool `json:"is_upgrade"`
	Banned               pgtype.Int8 `json:"banned"`
	Introduce            pgtype.Text `json:"introduce"`
}

type AccountStatus struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	StatusID  int64 `json:"status_id"`
	CreatedAt int64 `json:"created_at"`
}

type ApplyVoucher struct {
	ID        int64 `json:"id"`
	VoucherID int64 `json:"voucher_id"`
	OwnerID   int64 `json:"owner_id"`
}

type Follower struct {
	ID         int64  `json:"id"`
	FromFollow int64  `json:"from_follow"`
	ToFollow   int64  `json:"to_follow"`
	Status     string `json:"status"`
}

type Invoice struct {
	ID            int64              `json:"id"`
	FromAccountID int64              `json:"from_account_id"`
	ToAccountID   int64              `json:"to_account_id"`
	CreateAt      pgtype.Timestamptz `json:"create_at"`
	Status        pgtype.Text        `json:"status"`
	VoucherID     pgtype.Int8        `json:"voucher_id"`
	Total         pgtype.Numeric     `json:"total"`
}

type IssuePost struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	IsDeleted bool   `json:"is_deleted"`
}

type Locate struct {
	ID        int64       `json:"id"`
	AccountID int64       `json:"account_id"`
	Location  interface{} `json:"location"`
	Address   string      `json:"address"`
}

type Menu struct {
	ID        int64          `json:"id"`
	AccountID int64          `json:"account_id"`
	DishName  string         `json:"dish_name"`
	Quantity  int32          `json:"quantity"`
	Price     pgtype.Numeric `json:"price"`
	Img       string         `json:"img"`
	IsDelete  bool           `json:"is_delete"`
}

type Notification struct {
	ID           int64              `json:"id"`
	Message      string             `json:"message"`
	AccountID    int64              `json:"account_id"`
	TypeID       int32              `json:"type_id"`
	PostID       pgtype.Int8        `json:"post_id"`
	UserActionID int64              `json:"user_action_id"`
	InvoiceID    pgtype.Int8        `json:"invoice_id"`
	IsSeen       bool               `json:"is_seen"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
}

type NotificationType struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Permission struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Post struct {
	ID          int64              `json:"id"`
	PostTypeID  int32              `json:"post_type_id"`
	AccountID   int64              `json:"account_id"`
	PostTopID   pgtype.Int8        `json:"post_top_id"`
	Description pgtype.Text        `json:"description"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	Location    interface{}        `json:"location"`
	IsBanned    bool               `json:"is_banned"`
	IsDeleted   bool               `json:"is_deleted"`
}

type PostImage struct {
	ID       int64  `json:"id"`
	UrlImage string `json:"url_image"`
	PostID   int64  `json:"post_id"`
}

type PostType struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type ReactPost struct {
	AccountID int64 `json:"account_id"`
	PostID    int64 `json:"post_id"`
	State     int32 `json:"state"`
}

type ReportPost struct {
	ID        int64              `json:"id"`
	AccountID int64              `json:"account_id"`
	IssueID   int32              `json:"issue_id"`
	PostID    int64              `json:"post_id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

type ResetPassword struct {
	ID        pgtype.UUID        `json:"id"`
	UserID    int64              `json:"user_id"`
	ExpiresAt pgtype.Timestamptz `json:"expires_at"`
	IsActive  bool               `json:"is_active"`
}

type Role struct {
	ID   int32       `json:"id"`
	Name pgtype.Text `json:"name"`
}

type RolePermission struct {
	ID           int32 `json:"id"`
	PerID        int32 `json:"per_id"`
	RoleID       int32 `json:"role_id"`
	CanSelectAll bool  `json:"can_select_all"`
	CanSelect    bool  `json:"can_select"`
	CanInsert    bool  `json:"can_insert"`
	CanUpdate    bool  `json:"can_update"`
	CanDelete    bool  `json:"can_delete"`
	CanDoAll     bool  `json:"can_do_all"`
}

type Status struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID           int64              `json:"id"`
	Email        pgtype.Text        `json:"email"`
	Username     string             `json:"username"`
	HashPassword string             `json:"hash_password"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	IsDeleted    bool               `json:"is_deleted"`
}

type Voucher struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	ExpiredAt int64          `json:"expired_at"`
	IsDelete  bool           `json:"is_delete"`
	MinToUse  pgtype.Numeric `json:"min_to_use"`
	Discount  pgtype.Numeric `json:"discount"`
	MaxAmount pgtype.Numeric `json:"max_amount"`
}
