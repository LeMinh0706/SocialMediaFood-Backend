package service

import (
	"context"
	"fmt"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/models"
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

// Tam thoi chi lay account theo tu token, chua co thong tin khac
func (as *AccountService) GetAccountUser(ctx context.Context, user_id int64) ([]models.AccountResponse, error) {
	list, err := as.accountRepo.GetAccountByUserId(ctx, user_id)
	if err != nil {
		return []models.AccountResponse{}, err
	}
	res := models.ListAccountResponse(list)
	return res, nil
}

func (as *AccountService) GetAccountById(ctx context.Context, id int64) (models.AccountForPost, error) {
	var res models.AccountForPost
	account, err := as.accountRepo.GetAccountBydId(ctx, id)
	if err != nil {
		return res, err
	}
	res = models.AccountPost(account)
	return res, nil
}

func (as *AccountService) GetAccountForAction(ctx context.Context, user_id, id int64) (models.AccountForPost, error) {
	var res models.AccountForPost
	account, err := as.accountRepo.GetAccountBydId(ctx, id)
	if err != nil {
		return res, err
	}
	if account.UserID != user_id {
		return res, fmt.Errorf("not you")
	}
	res = models.AccountPost(account)
	return res, nil
}
