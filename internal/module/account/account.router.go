package account

import (
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewAccountRouter(r *gin.Engine, group *gin.RouterGroup, service IAccountService, token token.Maker) {
	ac, err := NewAccountController(service, token)
	if err != nil {
		log.Fatal(err)
	}
	accountGroup := r.Group(group.BasePath() + "/accounts")
	auth := accountGroup.Use(middlewares.AuthorizeMiddleware(token))
	{
		auth.GET("/me", ac.GetMe)
	}
}
