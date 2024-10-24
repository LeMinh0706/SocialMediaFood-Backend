package factory

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
)

type Factory struct {
	UserService *service.UserService
}

// Đang sửa lại thành cấu trúc cũ thì thành như này
func NewFactory() (*Factory, error) {
	//db connect
	pgx, err := db.GetDBConnection()
	if err != nil {
		return nil, err
	}

	//Repo
	queries := db.New(pgx)
	userRepo, err := repo.NewUserRepo(queries)
	if err != nil {
		return nil, err
	}

	//Service
	userService, err := service.NewUserService(userRepo)
	if err != nil {
		return nil, err
	}

	///return
	return &Factory{
		UserService: userService,
	}, nil
}
