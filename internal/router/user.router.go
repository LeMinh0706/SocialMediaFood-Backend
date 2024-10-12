package router

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/controller"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(r *gin.Engine, router *gin.RouterGroup, token token.Maker, userService *service.UserService, config util.Config) {
	userGroup := r.Group(router.BasePath() + "/accounts")
	auth := userGroup.Group("/").Use(middlewares.AuthorizeMiddleware(token))
	uc := controller.NewUserController(token, userService, config)
	{
		auth.GET("me", uc.GetMe)
		userGroup.GET(":id", uc.GetById)
		userGroup.POST("register", uc.Register)
		userGroup.POST("login", uc.Login)
	}
}
