package post

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewPostRouter(router *gin.RouterGroup, token token.Maker) {
	postGroup := router.Group("/post")
	{
		postGroup.POST("", NewPostController().CreatePost)
	}
}
