package router

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/controller"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewReactRouter(r *gin.Engine, router *gin.RouterGroup, token token.Maker, service *service.ReactPostService) {
	reactGroup := r.Group(router.BasePath() + "/react")
	auth := reactGroup.Group("").Use(middlewares.AuthorizeMiddleware(token))
	rc := controller.NewReactPostController(token, service)
	{
		auth.POST("", rc.LikePost)
		auth.DELETE("", rc.UnlikePost)
		reactGroup.GET(":id", rc.ListReactPost)
	}
}
