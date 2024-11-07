package account

import (
	"context"
	"fmt"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type AccountService struct {
	queries *db.Queries
}

// GetAccount implements IAccountService.
func (a *AccountService) GetAccount(ctx context.Context, id int64) (db.Account, error) {
	acc, err := a.queries.GetDetailAccount(ctx, id)
	if err != nil {
		return db.Account{}, err
	}
	return acc, nil
}

// GetAccountAction implements IAccountService.
func (a *AccountService) GetAccountAction(ctx context.Context, id int64, user_id int64) (db.GetAccountByIdRow, error) {
	var res db.GetAccountByIdRow
	acc, err := a.queries.GetAccountById(ctx, id)
	if err != nil {
		return res, err
	}
	if acc.ID != user_id {
		return res, fmt.Errorf("not you")
	}
	return acc, nil
}

// GetAccountById implements IAccountService.
func (a *AccountService) GetAccountById(ctx context.Context, id int64) (db.GetAccountByIdRow, error) {
	acc, err := a.queries.GetAccountById(ctx, id)
	if err != nil {
		return db.GetAccountByIdRow{}, err
	}
	return acc, nil
}

// GetAccountByUserId implements IAccountService.
func (a *AccountService) GetAccountByUserId(ctx context.Context, user_id int64) ([]AccountResponse, error) {
	res := make([]AccountResponse, 0)
	list, err := a.queries.GetAccountByUserId(ctx, user_id)
	if err != nil {
		return []AccountResponse{}, err
	}
	for _, element := range list {
		acc, _ := a.GetAccount(ctx, element)
		res = append(res, AccountRes(acc))
	}
	return res, nil
}

func NewAccountService(queries *db.Queries) IAccountService {
	return &AccountService{
		queries: queries,
	}
}
