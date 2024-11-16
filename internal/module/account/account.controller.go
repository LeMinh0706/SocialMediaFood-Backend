package account

import (
	"strconv"
	"strings"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	service IAccountService
	token   token.Maker
}

func NewAccountController(service IAccountService, token token.Maker) *AccountController {
	return &AccountController{
		service: service,
		token:   token,
	}
}

// Account godoc
// @Summary      Profile api
// @Description  To see the account, fetch profile
// @Tags         Accounts
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Security BearerAuth
// @Success      200  {object}  AccountResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /accounts/{id} [get]
func (ac *AccountController) GetAccount(g *gin.Context) {
	idStr := g.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrAccountID)
		return
	}
	acc, err := ac.service.GetAccount(g, id)
	if err != nil {
		response.ErrorResponse(g, 40414)
		return
	}
	response.SuccessResponse(g, 200, acc)
}

// Account godoc
// @Summary      It's you
// @Description  All your account is in here ->
// @Tags         Accounts
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Success      200  {object}  []AccountResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /accounts/me [get]
func (ac *AccountController) GetMe(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	me, err := ac.service.GetAccountByUserId(g, auth.UserId)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 200, me)
}

// Account godoc
// @Summary      Update Avatar
// @Description  Update Avatar
// @Tags         Accounts
// @Accept       multipart/form-data
// @Produce      json
// @Param        account_id formData string true "AccountID"
// @Param        image formData file true "Avatar Account"
// @Security BearerAuth
// @Success      200  {object}  AccountResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /accounts/avatar [put]
func (ac *AccountController) UpdateAvatar(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	accountStr := g.PostForm("account_id")
	account_id, err := strconv.ParseInt(accountStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrAccountID)
		return
	}
	var file string
	image, err := g.FormFile("image")
	if err == nil {
		var code int
		file, code = SaveAccountImage(g, "avatar", image)
		if code >= 40000 {
			response.ErrorResponse(g, code)
			return
		}
	}
	update, err := ac.service.UpdateAvatar(g, account_id, auth.UserId, file)
	if err != nil {
		response.ErrorNonKnow(g, 401, err.Error())
		return
	}
	response.SuccessResponse(g, 201, update)
}

// Account godoc
// @Summary      Update Background
// @Description  Update Background
// @Tags         Accounts
// @Accept       multipart/form-data
// @Produce      json
// @Param        account_id formData string true "AccountID"
// @Param        image formData file true "Background account"
// @Security BearerAuth
// @Success      200  {object}  AccountResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /accounts/background [put]
func (ac *AccountController) UpdateBackGround(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	accountStr := g.PostForm("account_id")
	account_id, err := strconv.ParseInt(accountStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrAccountID)
		return
	}
	var file string
	image, err := g.FormFile("image")
	if err == nil {
		var code int
		file, code = SaveAccountImage(g, "background", image)
		if code >= 40000 {
			response.ErrorResponse(g, code)
			return
		}
	}
	update, err := ac.service.UpdateBackground(g, account_id, auth.UserId, file)
	if err != nil {
		response.ErrorNonKnow(g, 401, err.Error())
		return
	}
	response.SuccessResponse(g, 201, update)
}

// Account godoc
// @Summary      Update Background
// @Description  Update Background
// @Tags         Accounts
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Param        request body UpdateNameReq true "request"
// @Security BearerAuth
// @Success      200  {object}  AccountResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /accounts/fullname/{id} [put]
func (ac *AccountController) UpdateName(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	accountStr := g.Param("id")
	var req UpdateNameReq
	account_id, err := strconv.ParseInt(accountStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrAccountID)
		return
	}
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
	}
	if strings.TrimSpace(req.Fullname) == "" {
		response.ErrorResponse(g, 40014)
		return
	}
	update, err := ac.service.UpdateName(g, account_id, auth.UserId, req.Fullname)
	if err != nil {
		response.ErrorNonKnow(g, 401, err.Error())
		return
	}
	response.SuccessResponse(g, 201, update)
}

func (as *AccountController) AddYourLocation(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	lng := g.PostForm("lng")
	lat := g.PostForm("lat")
	idStr := g.PostForm("account_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrAccountID)
		return
	}
	if !CheckValidPosition(g, lng, lat) {
		return
	}
	location, err := as.service.AddLocation(g, auth.UserId, id, lng, lat)
	if err != nil {
		response.ErrorNonKnow(g, 401, err.Error())
		return
	}
	response.SuccessResponse(g, 201, location)
}
