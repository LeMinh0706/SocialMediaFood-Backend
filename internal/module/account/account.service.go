package account

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type IAccountService interface {
	GetAccountByUserId(ctx context.Context, user_id int64) ([]AccountResponse, error)
	GetAccountById(ctx context.Context, id int64) (db.GetAccountByIdRow, error)
	GetAccount(ctx context.Context, id int64) (AccountResponse, error)
	GetAccountAction(ctx context.Context, id, user_id int64) (db.GetAccountByIdRow, error)
	UpdateName(ctx context.Context, id, user_id int64, name string) (AccountResponse, error)
	UpdateAvatar(ctx context.Context, id, user_id int64, url_avatar string) (AccountResponse, error)
	UpdateBackground(ctx context.Context, id, user_id int64, url_background string) (AccountResponse, error)
	AddLocation(ctx context.Context, user_id, account_id int64, lng, lat string) (db.CreateOwnerBranchRow, error)
	SearchingAccount(ctx context.Context, searching string, page, pageSize int32) ([]db.SearchingAccountsRow, error)
	AddEmail(ctx context.Context, id int64, email string) error
	Backup(ctx context.Context)
}
