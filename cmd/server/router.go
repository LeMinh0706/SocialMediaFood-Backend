package server

import (
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/factory"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/account"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/comment"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/follower"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/menu"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/notification"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/post"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/react"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/report"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/reset_password"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/user"
	"github.com/LeMinh0706/SocialMediaFood-Backend/swag/docs"
	swaggerfiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) NewRouter() {
	fac, err := factory.NewFactory(s.DBConn, s.Config)
	if err != nil {
		log.Fatal(err)
	}

	docs.SwaggerInfo.BasePath = "/api"
	a := s.Router.Group("/api")
	{
		s.Router.Use(middlewares.LoggerMiddleware(s.Logger))
		Static(s.Router)
		user.NewUserRouter(s.Router, a, fac.UserService, s.TokenMaker, s.RefeshMaker, s.Config)
		account.NewAccountRouter(s.Router, a, fac.AccountService, s.TokenMaker)
		post.NewPostRouter(s.Router, a, fac.PostService, s.TokenMaker)
		comment.NewCommentRouter(s.Router, a, fac.CommentService, s.TokenMaker)
		react.NewReactRouter(s.Router, a, fac.ReactService, s.TokenMaker)
		follower.NewFollowerRouter(s.Router, a, fac.FollowService, s.TokenMaker)
		report.NewReportRouter(s.Router, a, fac.ReportService, s.TokenMaker)
		reset_password.NewResetPasswordRouter(s.Router, a, fac.ResetPasswordService, s.Config, s.RefeshMaker)
		menu.NewMenuRouter(s.Router, a, fac.MenuService, s.TokenMaker)
		notification.NewNotificationRouter(s.Router, a, fac.NotificationService, s.TokenMaker)
	}

	s.Router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
