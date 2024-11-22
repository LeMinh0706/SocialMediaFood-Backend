package notification

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewNotificationRouter(r *gin.Engine, group *gin.RouterGroup, service INotificationService, token token.Maker) {
	n := NewNotificationController(service, token)
	auth := r.Group(group.BasePath() + "/notification").Use(middlewares.AuthorizeMiddleware(token))
	{
		auth.GET("/:id", n.GetYourNotification)
	}
}
