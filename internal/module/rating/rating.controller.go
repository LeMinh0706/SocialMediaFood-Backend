package rating

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type RatingController struct {
	service IRatingService
}

func NewRatingController(service IRatingService) *RatingController {
	return &RatingController{
		service: service,
	}
}

// Post godoc
// @Summary      Create rating
// @Description  Create rating
// @Tags         Rating
// @Accept       json
// @Produce      json
// @Param        request body RatingRequest true "request"
// @Security BearerAuth
// @Success      201  "Success"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /rating [post]
func (r *RatingController) CreateRating(g *gin.Context) {
	var req RatingRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	err := r.service.CreateRating(g, req)
	if err != nil {
		response.ErrorNonKnow(g, 50000, err.Error())
		return
	}
	response.SuccessResponse(g, 201, nil)
}
