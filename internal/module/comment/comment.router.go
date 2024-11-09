package comment

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewCommentRouter(r *gin.Engine, group *gin.RouterGroup, service ICommentService, token token.Maker) {
	cc := NewCommentController(service, token)

	commentGroup := r.Group(group.BasePath() + "/comments")
	auth := commentGroup.Group("").Use(middlewares.AuthorizeMiddleware(token))
	{
		auth.GET("", cc.GetListComment)
		auth.POST("", cc.CreateComment)
		auth.DELETE(":id", cc.DeleteComment)
		auth.PUT("", cc.UpdateComment)
	}
}
