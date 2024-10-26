package router

import (
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/controller"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewAccountRouter(r *gin.Engine, router *gin.RouterGroup, service *service.AccountService, token token.Maker) {
	ac, err := controller.NewAccountController(service, token)
	if err != nil {
		log.Fatal(err)
	}
	accountGroup := r.Group(router.BasePath() + "/accounts")
	auth := accountGroup.Group("").Use(middlewares.AuthorizeMiddleware(token))
	{
		auth.GET("/me", ac.GetMe)
	}
}
