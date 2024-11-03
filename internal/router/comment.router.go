package router

import (
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/controller"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewCommentRouter(r *gin.Engine, router *gin.RouterGroup, service *service.CommentService, token token.Maker) {
	cc, err := controller.NewCommentController(service)
	if err != nil {
		log.Fatal(err)
	}
	commentGroup := r.Group(router.BasePath() + "/comments")
	auth := commentGroup.Group("").Use(middlewares.AuthorizeMiddleware(token))
	{
		auth.POST("", cc.CreateComment)
		auth.PUT(":id", cc.UpdateComment)
		commentGroup.GET("", cc.GetListComment)
		auth.DELETE(":id", cc.DeteleComment)
	}
}
