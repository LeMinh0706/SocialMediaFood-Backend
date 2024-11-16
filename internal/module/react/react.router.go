package react

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewReactRouter(r *gin.Engine, group *gin.RouterGroup, service IReactService, token token.Maker) {
	rc := NewReactController(service, token)
	reactGroup := r.Group(group.BasePath() + "/react")
	auth := reactGroup.Group("").Use(middlewares.AuthorizeMiddleware(token))
	{
		auth.POST("", rc.CreateReact)
		auth.GET("/post/:id", rc.GetListReact)
		auth.PUT("", rc.ChangeReactState)
		auth.DELETE("", rc.UnReaction)
	}
}
