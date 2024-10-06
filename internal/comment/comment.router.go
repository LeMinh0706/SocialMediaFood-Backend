package comment

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewCommentRouter(router *gin.RouterGroup, token token.Maker) {
	commentGroup := router.Group("/comment")
	cc := NewCommentController(token)
	{
		// auth := commentGroup.Group("/").Use(middlewares.AuthorizeMiddleware(token))
		commentGroup.POST("", cc.CreateComment)
	}
}
