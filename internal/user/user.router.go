package user

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(r *gin.Engine, router *gin.RouterGroup, token token.Maker, config util.Config) {
	userGroup := r.Group(router.BasePath() + "/accounts")
	uc := NewUserController(token, config)
	{
		auth := userGroup.Group("/").Use(middlewares.AuthorizeMiddleware(token))
		auth.GET("me", uc.GetMe)
		userGroup.GET(":id", uc.GetById)
		userGroup.POST("register", uc.Register)
		userGroup.POST("login", uc.Login)
	}
}
