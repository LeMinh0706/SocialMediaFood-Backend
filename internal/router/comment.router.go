package router

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/controller"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewCommentRouter(r *gin.Engine, router *gin.RouterGroup, token token.Maker, commentService *service.CommentService) {
	commentGroup := r.Group(router.BasePath() + "/comment")
	cc := controller.NewCommentController(token, commentService)
	{
		auth := commentGroup.Group("/").Use(middlewares.AuthorizeMiddleware(token))
		auth.POST("", cc.CreateComment)
		auth.PUT(":id", cc.UpdateComment)
		auth.DELETE(":id", cc.DeleteComment)
		commentGroup.GET("", cc.ListComment)

	}
}
