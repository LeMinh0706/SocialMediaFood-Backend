package service

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
	"github.com/jackc/pgx/v5/pgtype"
)

type AccountService struct {
	accountRepo *repo.AccountRepo
}

func NewAccountService(repo *repo.AccountRepo) (*AccountService, error) {
	return &AccountService{
		accountRepo: repo,
	}, nil
}

func (as *AccountService) CreateAccount(ctx context.Context, user_id int64, fullname string, gender int32) (db.Account, error) {
	var res db.Account
	genderNull := pgtype.Int4{Int32: gender, Valid: true}
	account, err := as.accountRepo.CreateAccount(ctx, user_id, fullname, genderNull)
	if err != nil {
		return res, err
	}
	return account, nil
}

func (as *AccountService) GetAccountUser(ctx context.Context, user_id int64) ([]db.Account, error) {
	list, err := as.accountRepo.GetAccountByUserId(ctx, user_id)
	if err != nil {
		return []db.Account{}, err
	}
	return list, nil
}
