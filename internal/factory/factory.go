package factory

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
)

type Factory struct {
	UserRepo       *repo.UserRepository
	PostRepo       *repo.PostRepository
	CommentRepo    *repo.CommentRepository
	ReactRepo      *repo.ReactPostRepository
	UserService    *service.UserService
	PostService    *service.PostService
	CommentService *service.CommentService
	ReactService   *service.ReactPostService
}

// Đang sửa lại thành cấu trúc cũ thì thành như này
func NewFactory() (*Factory, error) {
	//db connect
	pg, err := db.GetDBConnection()
	if err != nil {
		return nil, err
	}

	queries := db.New(pg)

	//repository
	userRepo, err := repo.NewUserRepo(queries)
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

	reactRepo, err := repo.NewReactPostRepo(queries)
	if err != nil {
		return nil, err
	}

	//service
	userService := service.NewUserService(userRepo)
	postService := service.NewPostService(postRepo, userService)
	commentService := service.NewCommentService(commentRepo, userService, postService)
	reactService := service.NewReactPostService(reactRepo, userService, postService)

	///return
	return &Factory{
		UserRepo:       userRepo,
		PostRepo:       postRepo,
		CommentRepo:    commentRepo,
		ReactRepo:      reactRepo,
		UserService:    userService,
		PostService:    postService,
		CommentService: commentService,
		ReactService:   reactService,
	}, nil
}
