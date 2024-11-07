package router

import (
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/controller"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewFollowRouter(r *gin.Engine, router *gin.RouterGroup, service *service.FollowService, token token.Maker) {
	fc, err := controller.NewFollowController(service)
	if err != nil {
		log.Fatal(err)
	}
	followGroup := r.Group(router.BasePath() + "/follow")
	auth := followGroup.Use(middlewares.AuthorizeMiddleware(token))
	{
		auth.POST("", fc.CreateRequest)
		auth.GET("", fc.GetFollowStatus)
		auth.PUT("", fc.UpdateFriend)
		auth.DELETE("", fc.DeleteFollow)
		auth.GET("/self", fc.GetFollow)
		auth.GET("/other", fc.GetFollower)
		auth.GET("/friend", fc.GetFriend)
	}

}
