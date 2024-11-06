package account

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type AccountService struct {
	queries *db.Queries
}

// GetAccountAction implements IAccountService.
func (a *AccountService) GetAccountAction(ctx context.Context, id int64, user_id int64) (db.GetAccountByIdRow, error) {
	panic("unimplemented")
}

// GetAccountById implements IAccountService.
func (a *AccountService) GetAccountById(ctx context.Context, id int64) (db.GetAccountByIdRow, error) {
	panic("unimplemented")
}

// GetAccountByUserId implements IAccountService.
func (a *AccountService) GetAccountByUserId(ctx context.Context, user_id int64) ([]db.Account, error) {
	panic("unimplemented")
}

func NewAccountService(queries *db.Queries) IAccountService {
	return &AccountService{
		queries: queries,
	}
}
