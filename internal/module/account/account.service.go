package account

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type IAccountService interface {
	GetAccountByUserId(ctx context.Context, username string) (GetMeResponse, error)
	GetAccountById(ctx context.Context, id int64) (db.GetAccountByIdRow, error)
	GetAccount(ctx context.Context, id int64) (AccountResponse, error)
	GetAccountAction(ctx context.Context, id int64, username string) (db.Account, error)
	UpdateName(ctx context.Context, id int64, username, name string) (AccountResponse, error)
	UpdateAvatar(ctx context.Context, id int64, username, url_avatar string) (AccountResponse, error)
	UpdateBackground(ctx context.Context, id int64, username, url_background string) (AccountResponse, error)
	AddLocation(ctx context.Context, account_id int64, username, address, lng, lat string) (db.CreateOwnerBranchRow, error)
	SearchingAccount(ctx context.Context, searching string, page, pageSize int32) ([]db.SearchingAccountsRow, error)
	AddEmail(ctx context.Context, username, email string) error
	UpdateEmail(ctx context.Context, id int64, username, email string) error
	UpgradeOwnerRequest(ctx context.Context, id int64, username string) error
	GetUpgradePrice(ctx context.Context) (db.GetChoosePriceRow, error)
	CreateOwner(ctx context.Context, username string, arg CreateAccountVip) (AccountResponse, error)
}
