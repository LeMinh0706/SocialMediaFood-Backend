package router

import (
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/controller"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewPostRouter(r *gin.Engine, router *gin.RouterGroup, service *service.PostService, token token.Maker) {
	pc, err := controller.NewPostController(service)
	if err != nil {
		log.Fatal(err)
	}
	postGroup := r.Group(router.BasePath() + "/posts")
	auth := postGroup.Group("").Use(middlewares.AuthorizeMiddleware(token))
	{
		postGroup.GET("", pc.GetListPost)
		auth.GET("/person", pc.GetUserPost)
		// postGroup.GET(":id", pc.GetPost)
		auth.DELETE("/images/:id", pc.DeleteImagePost)
		auth.POST("", pc.CreatePost)
		auth.POST("/soft-delete/:id", pc.DeletePost)
		auth.PUT("", pc.UpdatePost)
	}
}
