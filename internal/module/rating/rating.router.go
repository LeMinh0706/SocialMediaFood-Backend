package rating

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewRatingRouter(r *gin.Engine, group *gin.RouterGroup, service IRatingService, token token.Maker) {
	rc := NewRatingController(service)
	ratingGroup := r.Group(group.BasePath() + "/rating").Use(middlewares.AuthorizeMiddleware(token))
	{
		ratingGroup.POST("", rc.CreateRating)
		ratingGroup.DELETE("", rc.DeleteRating)
		ratingGroup.GET("", rc.GetListRating)
	}
}
