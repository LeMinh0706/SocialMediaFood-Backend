package factory

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/account"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/comment"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/post"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/react"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Factory struct {
	UserService    user.IUserService
	AccountService account.IAccountService
	PostService    post.IPostService
	CommentService comment.ICommentService
	ReactService   react.IReactService
}

// Đang sửa lại thành cấu trúc cũ thì thành như này
func NewFactory(pq *pgxpool.Pool) (*Factory, error) {

	store := db.NewStore(pq)

	//Repo
	queries := db.New(pq)

	//Service

	userService := user.NewUserService(queries, store)
	accountService := account.NewAccountService(queries)
	postService := post.NewPostService(queries, accountService)
	commentService := comment.NewCommentService(queries, postService, accountService)
	reactService := react.NewReactService(queries, accountService, postService)
	///return
	return &Factory{
		UserService:    userService,
		AccountService: accountService,
		PostService:    postService,
		CommentService: commentService,
		ReactService:   reactService,
	}, nil
}
