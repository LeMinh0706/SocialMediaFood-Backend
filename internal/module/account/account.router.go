package account

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewAccountRouter(r *gin.Engine, group *gin.RouterGroup, service IAccountService, token token.Maker) {
	ac := NewAccountController(service, token)

	accountGroup := r.Group(group.BasePath() + "/accounts")
	auth := accountGroup.Group("").Use(middlewares.AuthorizeMiddleware(token))
	{
		auth.GET("/me", ac.GetMe)
		auth.GET("/:id", ac.GetAccount)
		auth.PUT("/avatar", ac.UpdateAvatar)
		auth.PUT("/background", ac.UpdateBackGround)
		auth.PUT("/fullname/:id", ac.UpdateName)
	}
}
