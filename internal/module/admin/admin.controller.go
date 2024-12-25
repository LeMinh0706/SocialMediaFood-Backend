package admin

import (
	"strconv"

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

// Admin godoc
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
	res, err := a.service.AddUpgragePrice(g, auth.Username, req.Title, req.Benefit, req.Price)
	if err != nil {
		handler.AdminErr(g, err)
		return
	}
	response.SuccessResponse(g, 201, res)
}

// Admin godoc
// @Summary      Only admin can do this
// @Description	 Get list price for upgrade in admin dashboard
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /admin/price [get]
func (a *AdminController) GetListUpgradePrice(g *gin.Context) {
	pageStr := g.Query("page")
	pageSizeStr := g.Query("page_size")
	page, pageSize := handler.CheckQuery(g, pageStr, pageSizeStr)
	if page == 0 || pageSize == 0 {
		return
	}
	list, err := a.service.GetUpgradePrice(g, page, pageSize)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 200, list)
}

// Admin godoc
// @Summary      Only admin can do this
// @Description	 Get list post were reported
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        account_id query int true "AccountID"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /admin/report [get]
func (a *AdminController) GetListReportPost(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	pageStr := g.Query("page")
	pageSizeStr := g.Query("page_size")
	page, pageSize := handler.CheckQuery(g, pageStr, pageSizeStr)
	if page == 0 || pageSize == 0 {
		return
	}
	accStr := g.Query("account_id")
	account_id, err := strconv.ParseInt(accStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 40004)
		return
	}
	res, err := a.service.GetListReportPost(g, auth.Username, account_id, page, pageSize)
	if err != nil {
		handler.AdminErr(g, err)
		return
	}
	response.SuccessResponse(g, 200, res)
}

// Admin godoc
// @Summary      Only admin can do this
// @Description	 Get list report, who report for upgrade in admin dashboard
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Param        account_id query int true "AccountID"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /admin/report/{id} [get]
func (a *AdminController) GetDetailReportPost(g *gin.Context) {
	idStr := g.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrBadRequestId)
		return
	}
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	pageStr := g.Query("page")
	pageSizeStr := g.Query("page_size")
	page, pageSize := handler.CheckQuery(g, pageStr, pageSizeStr)
	if page == 0 || pageSize == 0 {
		return
	}
	accStr := g.Query("account_id")
	account_id, err := strconv.ParseInt(accStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 40004)
		return
	}
	res, err := a.service.GetDetailReportPost(g, auth.Username, id, account_id, page, pageSize)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 200, res)
}

// Admin godoc
// @Summary      Only admin can do this
// @Description	 They're waiting for your accept
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /admin/upgrade-queue [get]
func (a *AdminController) GetUpgradeQueue(g *gin.Context) {
	pageStr := g.Query("page")
	pageSizeStr := g.Query("page_size")
	page, pageSize := handler.CheckQuery(g, pageStr, pageSizeStr)
	if page == 0 || pageSize == 0 {
		return
	}
	list, err := a.service.GetUpgradeQueue(g, page, pageSize)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 200, list)
}

// Admin godoc
// @Summary      Only admin can do this
// @Description	 They're waiting for your accept
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /admin/upgrade-queue/{id} [post]
func (a *AdminController) UpgradeSuccess(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	idStr := g.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrAccountID)
		return
	}
	acc, err := a.service.UpgradeSuccess(g, auth.Username, id)
	if err != nil {
		handler.AdminErr(g, err)
		return
	}
	response.SuccessResponse(g, 200, acc)
}

// Admin godoc
// @Summary      Only admin can do this
// @Description	 Cái này tạm thời chưa có, cảm ơn vì đã xem
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /admin/upgrade-queue/{id} [delete]
func (a *AdminController) UpgradeReject(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	idStr := g.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrAccountID)
		return
	}
	err = a.service.UpgradeReject(g, auth.Username, id)
	if err != nil {
		handler.AdminErr(g, err)
		return
	}
	response.SuccessResponse(g, 201, nil)
}

// Admin godoc
// @Summary      Only admin can do this
// @Description	 Cái này tạm thời chưa có, cảm ơn vì đã xem
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /admin/price-choosing/{id} [post]
func (a *AdminController) PriceChoosing(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	idStr := g.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrBadRequestId)
		return
	}
	err = a.service.ChoosingPrice(g, auth.Username, id)
	if err != nil {
		handler.AdminErr(g, err)
		return
	}
	response.SuccessResponse(g, 200, nil)
}
