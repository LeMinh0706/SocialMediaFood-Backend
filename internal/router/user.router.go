package router

import (
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/controller"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(r *gin.Engine, router *gin.RouterGroup, service *service.UserService, token token.Maker, config util.Config) {
	uc, err := controller.NewUserController(service, config, token)
	if err != nil {
		log.Fatal(err)
	}
	userGroup := r.Group(router.BasePath() + "/users")
	{
		userGroup.POST("/register", uc.Register)
		userGroup.POST("/login", uc.Login)
	}
}
