package router

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/controller"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewPostRouter(r *gin.Engine, router *gin.RouterGroup, token token.Maker, postService *service.PostService) {
	postGroup := r.Group(router.BasePath() + "/post")
	auth := postGroup.Group("").Use(middlewares.AuthorizeMiddleware(token))
	pc := controller.NewPostController(token, postService)
	{
		auth.POST("", pc.CreatePost)
		// postGroup.GET(":id", pc.GetPostById)
		postGroup.GET("", pc.GetListPost)
	}
}
