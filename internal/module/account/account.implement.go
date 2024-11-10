package account

import (
	"context"
	"fmt"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type AccountService struct {
	queries *db.Queries
}

// UpdateAvatar implements IAccountService.
func (a *AccountService) UpdateAvatar(ctx context.Context, id int64, user_id int64, url_avatar string) (AccountResponse, error) {
	panic("unimplemented")
}

// UpdateBackground implements IAccountService.
func (a *AccountService) UpdateBackground(ctx context.Context, id int64, user_id int64, url_background string) (AccountResponse, error) {
	panic("unimplemented")
}

// UpdateName implements IAccountService.
func (a *AccountService) UpdateName(ctx context.Context, id int64, user_id int64, name string) (AccountResponse, error) {
	panic("unimplemented")
}

// Backup implements IAccountService.
func (a *AccountService) Backup(ctx context.Context) {
	panic("unimplemented")
}

// GetAccount implements IAccountService.
func (a *AccountService) GetAccount(ctx context.Context, id int64) (AccountResponse, error) {
	acc, err := a.queries.GetDetailAccount(ctx, id)
	if err != nil {
		return AccountResponse{}, err
	}
	res := AccountRes(acc)
	return res, nil
}

// GetAccountAction implements IAccountService.
func (a *AccountService) GetAccountAction(ctx context.Context, id int64, user_id int64) (db.GetAccountByIdRow, error) {
	var res db.GetAccountByIdRow
	acc, err := a.queries.GetAccountById(ctx, id)
	if err != nil {
		return res, err
	}
	if acc.UserID != user_id {
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
		res = append(res, acc)
	}
	return res, nil
}

func NewAccountService(queries *db.Queries) IAccountService {
	return &AccountService{
		queries: queries,
	}
}
