package account

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type IAccountService interface {
	GetAccountByUserId(ctx context.Context, user_id int64) ([]db.Account, error)
	GetAccountById(ctx context.Context, id int64) (db.GetAccountByIdRow, error)
	GetAccountAction(ctx context.Context, id, user_id int64) (db.GetAccountByIdRow, error)
}
