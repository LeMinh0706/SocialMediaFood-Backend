package routers

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/controllers"
	"github.com/gin-gonic/gin"
)

func NewPostRouter(router *gin.RouterGroup) {
	postGroup := router.Group("/post")
	{
		postGroup.POST("", controllers.NewPostController().CreatePost)
	}
}
