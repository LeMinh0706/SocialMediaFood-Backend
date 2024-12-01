package admin

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/handler"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	service IAdminService
}

func NewAdminController(service IAdminService) AdminController {
	return AdminController{
		service: service,
	}
}

// Account godoc
// @Summary      Only admin can do this
// @Description	 Add price for upgrade
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        request body AddUpgradePrice true "request"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /admin/price [post]
func (a *AdminController) AddUpgradePrice(g *gin.Context) {
	var req AddUpgradePrice
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, response.ErrBadRequest)
		return
	}
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	res, err := a.service.AddUpgragePrice(g, auth.UserId, req.Price)
	if err != nil {
		handler.AdminErr(g, err)
		return
	}
	response.SuccessResponse(g, 201, res)
}
