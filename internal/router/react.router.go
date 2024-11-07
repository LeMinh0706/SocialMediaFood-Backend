package router

import (
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/controller"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewReactRouter(r *gin.Engine, router *gin.RouterGroup, service *service.ReactService, token token.Maker) {
	rc, err := controller.NewReactController(service)
	if err != nil {
		log.Fatal(err)
	}
	reactGroup := r.Group(router.BasePath() + "/react")
	auth := reactGroup.Use(middlewares.AuthorizeMiddleware(token))
	{
		auth.POST("", rc.CreateReact)
		auth.GET("", rc.GetReactPost)
		auth.PUT("", rc.UpdateReact)
		auth.DELETE("", rc.UnlikePost)
	}
}
