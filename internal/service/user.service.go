package service

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
)

type UserService struct {
	userRepo *repo.UserRepository
}

func NewUserService(repo *repo.UserRepository) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

func (us *UserService) Register(ctx context.Context, username, password, fullname string, emailReq string, gender int32) (db.User, error) {
	hashPassword, err := util.HashPashword(password)
	if err != nil {
		return db.User{}, err
	}
	if strings.TrimSpace(fullname) == "" {
		fullname = username
	}
	var email sql.NullString
	if strings.TrimSpace(emailReq) == "" {
		email = sql.NullString{Valid: false}
	} else {
		email = sql.NullString{String: emailReq, Valid: true}
	}

	user, err := us.userRepo.CreateUser(ctx, username, hashPassword, fullname, email, 3, gender)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}

func (us *UserService) GetMe(ctx context.Context, username string) (response.UserResponse, error) {
	user, err := us.userRepo.GetUser(ctx, username)
	if err != nil {
		return response.UserRes(user), err
	}
	res := response.UserRes(user)
	return res, nil
}

func (us *UserService) Login(ctx context.Context, username, password string) (response.UserResponse, error) {
	user, err := us.userRepo.GetUser(ctx, username)

	if err != nil {
		return response.UserRes(user), err
	}

	if err := util.CheckPassword(password, user.HashPashword); err != nil {
		return response.UserRes(user), err
	}
	res := response.UserRes(user)
	return res, nil
}

func (us *UserService) GetUser(ctx context.Context, id int64) (db.GetUserByIdRow, error) {
	user, err := us.userRepo.GetUserById(ctx, id)

	if err == sql.ErrNoRows {
		return db.GetUserByIdRow{}, fmt.Errorf("User does not exist")
	}
	return user, nil
}
