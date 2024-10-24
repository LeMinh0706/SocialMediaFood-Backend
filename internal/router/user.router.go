package router

import (
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/controller"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(r *gin.Engine, router *gin.RouterGroup, service *service.UserService) {
	userGroup := r.Group(router.BasePath() + "/users")
	uc, err := controller.NewUserController(service)
	if err != nil {
		log.Fatal(err)
		return
	}
	{
		userGroup.POST("/register", uc.Register)
	}
}
