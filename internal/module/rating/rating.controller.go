package rating

import (
	"strconv"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/handler"
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

// Raing godoc
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

// Rating godoc
// @Summary      Delete rating
// @Description  Delete rating
// @Tags         Rating
// @Accept       json
// @Produce      json
// @Param        request body DeleteRatingRequest true "request"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /rating [delete]
func (r *RatingController) DeleteRating(g *gin.Context) {
	var req DeleteRatingRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	err := r.service.DeleteRating(g, req.FromAccountID, req.ToAccountID)
	if err != nil {
		response.ErrorNonKnow(g, 50000, err.Error())
		return
	}
	response.SuccessResponse(g, 204, nil)
}

// Post godoc
// @Summary      Get list rating
// @Description  Get list rating
// @Tags         Rating
// @Accept       json
// @Produce      json
// @Param        account_id query int true "Your account id"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Security BearerAuth
// @Success      200  {object}  []ListRating
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /rating [get]
func (r *RatingController) GetListRating(g *gin.Context) {
	fromStr := g.Query("account_id")
	from, err := strconv.ParseInt(fromStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 40004)
		return
	}
	pageStr := g.Query("page")
	pageSizeStr := g.Query("page_size")
	page, pageSize := handler.CheckQuery(g, pageStr, pageSizeStr)
	if page == 0 || pageSize == 0 {
		return
	}
	list, err := r.service.GetListRating(g, from, page, pageSize)
	if err != nil {
		response.ErrorNonKnow(g, 50000, err.Error())
		return
	}
	response.SuccessResponse(g, 200, list)
}
