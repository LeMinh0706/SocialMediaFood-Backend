package router

import (
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/controller"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/gin-gonic/gin"
)

func NewCommentRouter(r *gin.Engine, router *gin.RouterGroup, service *service.CommentService) {
	cc, err := controller.NewCommentController(service)
	if err != nil {
		log.Fatal(err)
	}
	commentGroup := r.Group(router.BasePath() + "/comments")
	{
		commentGroup.POST("", cc.CreateComment)
	}
}
