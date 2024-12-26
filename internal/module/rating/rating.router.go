package rating

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewRatingRouter(r *gin.Engine, group *gin.RouterGroup, service IRatingService, token token.Maker) {
	rc := NewRatingController(service)
	ratingGroup := r.Group(group.BasePath() + "/rating")
	{
		ratingGroup.POST("", rc.CreateRating)
	}
}
