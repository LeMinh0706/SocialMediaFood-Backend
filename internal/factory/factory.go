package factory

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/account"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/admin"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/comment"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/follower"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/menu"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/notification"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/post"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/rating"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/react"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/report"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/reset_password"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/user"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Factory struct {
	UserService          user.IUserService
	AccountService       account.IAccountService
	PostService          post.IPostService
	CommentService       comment.ICommentService
	ReactService         react.IReactService
	FollowService        follower.IFollowerService
	ReportService        report.IReportService
	ResetPasswordService reset_password.IResetPasswordService
	MenuService          menu.IMenuService
	NotificationService  notification.INotificationService
	AdminService         admin.IAdminService
	RatingService        rating.IRatingService
}

// Đang sửa lại thành cấu trúc cũ thì thành như này
func NewFactory(pq *pgxpool.Pool, config util.Config) (*Factory, error) {

	//Store transaction
	store := db.NewStore(pq)

	//Repo
	q := db.New(pq)

	//Service

	userService := user.NewUserService(q, store)
	accountService := account.NewAccountService(q)
	postService := post.NewPostService(q, store, accountService)
	notificationService := notification.NewNotificationService(q, accountService)
	commentService := comment.NewCommentService(q, postService, accountService, notificationService)
	reactService := react.NewReactService(q, accountService, postService, notificationService)
	followService := follower.NewFollowerService(q, accountService)
	reportService := report.NewReportService(q, accountService)
	resetService := reset_password.NewResetPasswordService(q)
	menuService := menu.NewMenuService(q)
	adminService := admin.NewAdminService(q, accountService, postService)
	ratingService := rating.NewRatingService(q, accountService)
	///return
	return &Factory{
		UserService:          userService,
		AccountService:       accountService,
		PostService:          postService,
		CommentService:       commentService,
		ReactService:         reactService,
		FollowService:        followService,
		ReportService:        reportService,
		ResetPasswordService: resetService,
		MenuService:          menuService,
		NotificationService:  notificationService,
		AdminService:         adminService,
		RatingService:        ratingService,
	}, nil
}
