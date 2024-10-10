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
	UserService    *service.UserService
	PostService    *service.PostService
	CommentService *service.CommentService
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

	//service
	userService := service.NewUserService(userRepo)
	postService := service.NewPostService(postRepo, userService)
	commentService := service.NewCommentService(commentRepo, userService, postService)

	return &Factory{
		UserRepo:       userRepo,
		PostRepo:       postRepo,
		CommentRepo:    commentRepo,
		UserService:    userService,
		PostService:    postService,
		CommentService: commentService,
	}, nil
}
