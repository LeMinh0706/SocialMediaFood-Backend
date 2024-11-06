package user

import (
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(r *gin.Engine, router *gin.RouterGroup, service IUserService, token token.Maker, config util.Config) {
	uc, err := NewUserController(service, config, token)
	if err != nil {
		log.Fatal(err)
	}
	userGroup := r.Group(router.BasePath() + "/user")
	{
		userGroup.POST("/login", uc.Login)
		userGroup.POST("/register", uc.RegisterTx)
	}
}
