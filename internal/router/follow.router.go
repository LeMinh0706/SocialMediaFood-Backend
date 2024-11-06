package router

import (
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/controller"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/gin-gonic/gin"
)

func NewFollowRouter(r *gin.Engine, router *gin.RouterGroup, service *service.FollowService) {
	fc, err := controller.NewFollowController(service)
	if err != nil {
		log.Fatal(err)
	}
	followGroup := r.Group(router.BasePath() + "/follow")
	{
		followGroup.POST("", fc.CreateRequest)
		followGroup.GET("", fc.GetFollowStatus)
		followGroup.PUT("", fc.UpdateFriend)
		followGroup.DELETE("", fc.DeleteFollow)
	}

}
