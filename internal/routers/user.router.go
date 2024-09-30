package routers

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/controllers"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/token"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(router *gin.RouterGroup, token token.Maker) {
	userGroup := router.Group("/accounts")
	uc := controllers.NewUserController(token)
	{
		userGroup.GET(":id", uc.GetById)
		userGroup.POST("register", uc.Register)
		userGroup.POST("login", uc.Login)
	}
}
