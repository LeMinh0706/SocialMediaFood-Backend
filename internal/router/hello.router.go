package router

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewHelloRouter(router *gin.RouterGroup) {
	userGroup := router.Group("/hello")
	uc := controller.NewHelloController()
	{
		userGroup.GET(":name", uc.GetHelloParam)
		userGroup.GET("", uc.GetHelloQuery)
	}
}
