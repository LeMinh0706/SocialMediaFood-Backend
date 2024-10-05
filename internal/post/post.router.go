package post

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewPostRouter(router *gin.RouterGroup, token token.Maker) {
	postGroup := router.Group("/post")
	pc := NewPostController(token)
	{
		auth := postGroup.Group("/").Use(middlewares.AuthorizeMiddleware(token))
		auth.POST("", pc.CreatePost)
		postGroup.GET(":id", pc.GetPostById)
		postGroup.GET("", pc.GetListPost)
	}
}
