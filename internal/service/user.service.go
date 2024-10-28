package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/jackc/pgx/v5"
)

type UserService struct {
	userRepo       *repo.UserRepo
	accountService *AccountService
}

func NewUserService(repo *repo.UserRepo, accountService *AccountService) (*UserService, error) {
	return &UserService{
		userRepo:       repo,
		accountService: accountService,
	}, nil
}

func (us *UserService) Login(ctx context.Context, username, password string) (db.LoginRow, error) {
	var res db.LoginRow
	user, err := us.userRepo.Login(ctx, username)
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

func (us *UserService) RegisterTx(ctx context.Context, req db.RegisterRequest) (db.RegisterRow, error) {
	var res db.RegisterRow
	if !util.UsernameNotSpace(req.Username) {
		return res, fmt.Errorf("username don't have a space")
	}
	if strings.TrimSpace(req.Email) != "" {
		if !util.EmailCheck(req.Email) {
			return res, fmt.Errorf("this mail is invalid")
		}
	}
	user, err := us.userRepo.RegisterTx(ctx, req)
	if err != nil {
		return res, err
	}
	return user, nil
}
