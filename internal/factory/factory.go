package factory

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Factory struct {
	UserService    *service.UserService
	AccountService *service.AccountService
}

// Đang sửa lại thành cấu trúc cũ thì thành như này
func NewFactory(pq *pgxpool.Pool) (*Factory, error) {

	store := db.NewStore(pq)

	//Repo
	queries := db.New(pq)
	userRepo, err := repo.NewUserRepo(queries, store)
	if err != nil {
		return nil, err
	}

	accountRepo, err := repo.NewAccountRepo(queries)
	if err != nil {
		return nil, err
	}

	//Service
	accountService, err := service.NewAccountService(accountRepo)
	if err != nil {
		return nil, err
	}
	userService, err := service.NewUserService(userRepo, accountService)
	if err != nil {
		return nil, err
	}

	///return
	return &Factory{
		UserService:    userService,
		AccountService: accountService,
	}, nil
}
