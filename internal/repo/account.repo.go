package repo

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/jackc/pgx/v5/pgtype"
)

type AccountRepo struct {
	queries *db.Queries
}

func NewAccountRepo(queries *db.Queries) (*AccountRepo, error) {
	return &AccountRepo{
		queries: queries,
	}, nil
}

func (ar *AccountRepo) CreateAccount(ctx context.Context, user_id int64, fullname string, gender pgtype.Int4) (db.Account, error) {
	return ar.queries.CreateAccounts(ctx, db.CreateAccountsParams{
		UserID:               user_id,
		Fullname:             fullname,
		Gender:               gender,
		UrlAvatar:            util.RandomAvatar(gender.Int32),
		UrlBackgroundProfile: db.GetBackground(),
	})
}

func (ar *AccountRepo) GetAccountByUserId(ctx context.Context, user_id int64) ([]db.GetAccountByUserIdRow, error) {
	return ar.queries.GetAccountByUserId(ctx, user_id)
}

func (ar *AccountRepo) GetAccountBydId(ctx context.Context, id int64) (db.Account, error) {
	return ar.queries.GetAccountById(ctx, id)
}
