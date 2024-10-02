package services

import (
	"context"
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
)

type UserService struct {
	userRepo *repo.UserRepository
}

func NewUserService() *UserService {
	repo, err := repo.NewUserRepo()
	if err != nil {
		log.Fatal("Error:", err)
	}
	return &UserService{
		userRepo: repo,
	}
}

func (us *UserService) Register(ctx context.Context, username, password string) (db.User, error) {
	hashPassword, err := util.HashPashword(password)
	if err != nil {
		return db.User{}, err
	}
	user, err := us.userRepo.CreateUser(ctx, username, hashPassword, 3)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}

func (us *UserService) GetMe(ctx context.Context, username string) (response.UserResponse, error) {
	user, err := us.userRepo.GetUser(ctx, username)
	if err != nil {
		return response.UserResponse{}, err
	}
	res := response.UserResponse{ID: user.ID, Fullname: user.Fullname, Gender: user.Gender, RoleID: user.RoleID, DateCreateAccount: user.DateCreateAccount}
	return res, nil
}

func (us *UserService) Login(ctx context.Context, username, password string) (db.User, error) {
	user, err := us.userRepo.GetUser(ctx, username)

	if err != nil {
		return db.User{}, err
	}

	if err := util.CheckPassword(password, user.HashPashword); err != nil {
		return db.User{}, err
	}

	return user, nil
}

func (us *UserService) GetUser(ctx context.Context, id int64) (db.GetUserByIdRow, error) {
	user, err := us.userRepo.GetUserById(ctx, id)

	if err != nil {
		return db.GetUserByIdRow{}, err
	}
	return user, nil
}
