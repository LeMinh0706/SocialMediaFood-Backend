package service

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService(repo *repo.UserRepo) (*UserService, error) {
	return &UserService{
		userRepo: repo,
	}, nil
}

func (us *UserService) Register(ctx context.Context, username, password string, email pgtype.Text) (db.RegisterRow, error) {
	var res db.RegisterRow
	hash, err := util.HashPashword(password)
	if err != nil {
		return res, err
	}
	user, err := us.userRepo.Register(ctx, username, hash, email)
	if err != nil {
		return res, err
	}
	return user, nil
}
