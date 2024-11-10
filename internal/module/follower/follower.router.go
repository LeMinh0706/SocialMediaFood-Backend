package follower

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewFollowerRouter(r *gin.Engine, group *gin.RouterGroup, service IFollowerService, token token.Maker) {
	fc := NewFollowerController(service, token)
	followGroup := r.Group(group.BasePath() + "/follower")
	auth := followGroup.Group("").Use(middlewares.AuthorizeMiddleware(token))
	{
		auth.POST("", fc.FollowRequest)
		auth.GET("/status", fc.GetFollowStatus)
		auth.GET("", fc.GetFollowType)
		auth.PUT("", fc.UpdateFriend)
		auth.DELETE("", fc.UnFollow)
	}
}
