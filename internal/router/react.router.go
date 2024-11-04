package router

import (
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/controller"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/gin-gonic/gin"
)

func NewReactRouter(r *gin.Engine, router *gin.RouterGroup, service *service.ReactService) {
	rc, err := controller.NewReactController(service)
	if err != nil {
		log.Fatal(err)
	}
	reactGroup := r.Group(router.BasePath() + "/react")
	{
		reactGroup.POST("", rc.CreateReact)
		reactGroup.GET("", rc.GetReactPost)
		reactGroup.PUT("", rc.UpdateReact)
		reactGroup.DELETE("", rc.UnlikePost)
	}
}
