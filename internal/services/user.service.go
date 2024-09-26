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
	user, err := us.userRepo.CreateUser(username, password, 3)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}
