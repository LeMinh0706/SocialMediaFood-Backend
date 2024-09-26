package routers

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/controllers"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(router *gin.RouterGroup) {
	userGroup := router.Group("/account")
	{
		userGroup.POST("", controllers.NewUserController().Register)
	}
}
