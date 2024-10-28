package service

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
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

// Tam thoi chi lay account theo tu token, chua co thong tin khac 24/10
func (as *AccountService) GetAccountUser(ctx context.Context, user_id int64) ([]response.AccountResponse, error) {
	list, err := as.accountRepo.GetAccountByUserId(ctx, user_id)
	if err != nil {
		return []response.AccountResponse{}, err
	}
	res := response.ListAccountResponse(list)
	return res, nil
}

func (as *AccountService) GetAccountById(ctx context.Context, id int64) (db.Account, error) {
	var res db.Account
	account, err := as.accountRepo.GetAccountBydId(ctx, id)
	if err != nil {
		return res, err
	}
	return account, nil
}
