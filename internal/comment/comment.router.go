package comment

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewCommentRouter(r *gin.Engine, router *gin.RouterGroup, token token.Maker) {
	commentGroup := r.Group(router.BasePath() + "/comment")
	cc := NewCommentController(token)
	{
		auth := commentGroup.Group("/").Use(middlewares.AuthorizeMiddleware(token))
		auth.POST("", cc.CreateComment)
		auth.PUT(":id", cc.UpdateComment)
		auth.DELETE(":id", cc.DeleteComment)
		commentGroup.GET("", cc.ListComment)

	}
}
