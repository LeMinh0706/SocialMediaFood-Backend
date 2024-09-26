package services

import (
	"context"
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
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
	user, err := us.userRepo.CreateUser(ctx, username, password, 3)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}

func (us *UserService) UserExist(ctx context.Context, username string) error {
	_, err := us.userRepo.UserExist(ctx, username)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) GetUser(ctx context.Context, id int64) (db.GetUserRow, error) {
	user, err := us.userRepo.GetUser(ctx, id)
	if err != nil {
		return db.GetUserRow{}, err
	}
	return user, nil
}
