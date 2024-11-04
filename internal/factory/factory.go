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
	PostService    *service.PostService
	CommentService *service.CommentService
	ReactService   *service.ReactService
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

	postRepo, err := repo.NewPostRepo(queries)
	if err != nil {
		return nil, err
	}

	commentRepo, err := repo.NewCommentRepo(queries)
	if err != nil {
		return nil, err
	}

	reactRepo, err := repo.NewReactRepo(queries)

	//Service
	accountService, err := service.NewAccountService(accountRepo)
	if err != nil {
		return nil, err
	}
	userService, err := service.NewUserService(userRepo, accountService)
	if err != nil {
		return nil, err
	}
	postService, err := service.NewPostService(postRepo, accountService)
	if err != nil {
		return nil, err
	}
	commentSerice, err := service.NewCommentService(commentRepo, accountService, postService)
	if err != nil {
		return nil, err
	}
	reactService, err := service.NewReactService(reactRepo)
	if err != nil {
		return nil, err
	}

	///return
	return &Factory{
		UserService:    userService,
		AccountService: accountService,
		PostService:    postService,
		CommentService: commentSerice,
		ReactService:   reactService,
	}, nil
}
