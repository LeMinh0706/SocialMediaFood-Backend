package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
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

func (us *UserService) Register(ctx context.Context, req response.RegisterRequest) (response.RegisterResponse, error) {
	var res response.RegisterResponse
	hash, err := util.HashPassword(req.Password)
	if err != nil {
		return res, err
	}
	var nullmail pgtype.Text
	if strings.TrimSpace(req.Email) == "" {
		nullmail = pgtype.Text{Valid: false}
	} else {
		nullmail = pgtype.Text{String: req.Email, Valid: true}
	}
	user, err := us.userRepo.Register(ctx, req.Username, hash, nullmail)
	if err != nil {
		return res, err
	}
	us.accountService.CreateAccount(ctx, user.ID, req.Fullname, req.Gender)
	res = response.RegisterRes(user)
	return res, nil
}

func (us *UserService) Login(ctx context.Context, username, password string) (db.LoginRow, error) {
	var res db.LoginRow
	user, err := us.userRepo.Login(ctx, username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return res, fmt.Errorf("wrong username or password")
		}
		return res, err
	}
	if err = util.CheckPassword(password, user.HashPassword); err != nil {
		return res, fmt.Errorf("wrong username or password %v", password)
	}
	return user, nil
}
