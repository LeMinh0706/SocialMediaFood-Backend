package user

import (
	"context"
	"fmt"
	"strings"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/jackc/pgx/v5"
)

type UserService struct {
	queries *db.Queries
	store   *db.Store
}

// Backup implements IUserService.
func (u *UserService) Backup(ctx context.Context) {
	panic("unimplemented")
}

// Login implements IUserService.
func (u *UserService) Login(ctx context.Context, username string, password string) (db.LoginRow, error) {
	var res db.LoginRow
	user, err := u.queries.Login(ctx, username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return res, fmt.Errorf("wrong username")
		}
		return res, err
	}
	if err = util.CheckPassword(password, user.HashPassword); err != nil {
		return res, fmt.Errorf("wrong password")
	}
	return user, nil
}

// Register implements IUserService.
func (u *UserService) Register(ctx context.Context, req db.RegisterRequest) (db.RegisterRow, error) {
	var res db.RegisterRow
	if !util.UsernameNotSpace(req.Username) {
		return res, fmt.Errorf("username don't have a space")
	}
	if strings.TrimSpace(req.Email) != "" {
		if !util.EmailCheck(req.Email) {
			return res, fmt.Errorf("this mail is invalid")
		}
	}
	user, err := u.store.CreateAccountTx(ctx, req)
	if err != nil {
		return res, err
	}
	return user, nil
}

func NewUserService(queries *db.Queries, store *db.Store) IUserService {
	return &UserService{
		queries: queries,
		store:   store,
	}
}
