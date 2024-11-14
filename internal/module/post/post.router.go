package post

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewPostRouter(r *gin.Engine, group *gin.RouterGroup, service IPostService, token token.Maker) {
	pc := NewPostController(service, token)

	postGroup := r.Group(group.BasePath() + "/posts")
	auth := postGroup.Group("").Use(middlewares.AuthorizeMiddleware(token))
	{
		postGroup.GET("", pc.GetListPost)
		auth.POST("", pc.CreatePost)
		auth.GET("/person", pc.GetPersonPost)
		auth.PUT("/:id", pc.UpdateContentPost)
		auth.DELETE("/images/:id", pc.DeleteImage)
		auth.POST("/soft-delete/:id", pc.DeletePost)
		postGroup.GET("/home", pc.GetHomePagePost)
	}
}
