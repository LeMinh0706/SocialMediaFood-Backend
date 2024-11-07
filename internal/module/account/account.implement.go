package account

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type IAccountService interface {
	GetAccountByUserId(ctx context.Context, user_id int64) ([]AccountResponse, error)
	GetAccountById(ctx context.Context, id int64) (db.GetAccountByIdRow, error)
	GetAccount(ctx context.Context, id int64) (db.Account, error)
	GetAccountAction(ctx context.Context, id, user_id int64) (db.GetAccountByIdRow, error)
	UpdateName(ctx context.Context, id, user_id int64, name string) (AccountResponse, error)
	UpdateAvatar(ctx context.Context, id, user_id int64, url_avatar string) (AccountResponse, error)
	UpdateBackground(ctx context.Context, id, user_id int64, url_background string) (AccountResponse, error)
	Backup(ctx context.Context)
}
