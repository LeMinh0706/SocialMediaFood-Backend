package menu

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewMenuRouter(r *gin.Engine, group *gin.RouterGroup, service IMenuService, token token.Maker) {
	mc := NewMenuController(service, token)

	menuGroup := r.Group(group.BasePath() + "/menu").Use(middlewares.AuthorizeMiddleware(token))
	{
		menuGroup.POST("", mc.CreateNewFood)
		menuGroup.GET(":id", mc.GetMenu)
	}
}
