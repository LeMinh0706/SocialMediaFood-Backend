package user

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(r *gin.Engine, router *gin.RouterGroup, service IUserService, token token.Maker, config util.Config) {
	uc := NewUserController(service, config, token)

	userGroup := r.Group(router.BasePath() + "/users")
	{
		userGroup.POST("/login", uc.Login)
		userGroup.POST("/register", uc.RegisterTx)
	}
}
