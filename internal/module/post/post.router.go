package post

import (
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewPostRouter(r *gin.Engine, group *gin.RouterGroup, service IPostService, token token.Maker) {
	pc, err := NewPostController(service, token)
	if err != nil {
		log.Fatal(err)
	}
	postGroup := r.Group(group.BasePath() + "/posts")
	auth := postGroup.Group("").Use(middlewares.AuthorizeMiddleware(token))
	{
		postGroup.GET("", pc.GetListPost)
		auth.POST("", pc.CreatePost)
		auth.GET("/person", pc.GetPersonPost)
	}
}
