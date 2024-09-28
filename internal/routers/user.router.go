package routers

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/controllers"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(router *gin.RouterGroup) {
	userGroup := router.Group("/accounts")
	{
		userGroup.GET(":id", controllers.NewUserController().GetById)
		userGroup.POST("register", controllers.NewUserController().Register)
		userGroup.GET("login", controllers.NewUserController().Login)
	}
}
