package account

import (
	"strconv"
	"strings"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/handler"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
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
// @Success      201  {object}  AccountResponse
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
		file, code = handler.SaveImage(g, "avatar", image)
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
// @Success      201  {object}  AccountResponse
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
		file, code = handler.SaveImage(g, "background", image)
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
// @Summary      Update fullname
// @Description  Update fullname
// @Tags         Accounts
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Param        request body UpdateNameReq true "request"
// @Security BearerAuth
// @Success      201  {object}  AccountResponse
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
	address := g.PostForm("address")
	idStr := g.PostForm("account_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrAccountID)
		return
	}
	if !handler.CheckValidPosition(g, lng, lat) {
		return
	}
	location, err := as.service.AddLocation(g, auth.UserId, id, address, lng, lat)
	if err != nil {
		response.ErrorNonKnow(g, 401, err.Error())
		return
	}
	response.SuccessResponse(g, 201, location)
}

// Account godoc
// @Summary      Profile api
// @Description  To see the account, searching account
// @Tags         Accounts
// @Accept       json
// @Produce      json
// @Param        name query string true "Who you want to search"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Security BearerAuth
// @Success      200  {object}  []db.SearchingAccountsRow
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /accounts/searching [get]
func (as *AccountController) Searching(g *gin.Context) {
	param := g.Query("name")
	if strings.TrimSpace(param) == "" {
		response.ErrorResponse(g, response.ErrInputSearch)
		return
	}
	pageStr := g.Query("page")
	pageSizeStr := g.Query("page_size")
	page, pageSize := handler.CheckQuery(g, pageStr, pageSizeStr)
	if page == 0 || pageSize == 0 {
		return
	}
	result, err := as.service.SearchingAccount(g, param, page, pageSize)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 200, result)
}

// Account godoc
// @Summary      Add your email
// @Description  Add your email, you can also update email here
// @Tags         Accounts
// @Accept       json
// @Produce      json
// @Param        request body EmailRequest true "request"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /accounts [put]
func (as *AccountController) AddEmail(g *gin.Context) {
	var req EmailRequest
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, response.ErrEmailInvalid)
		return
	}
	err := as.service.AddEmail(g, auth.UserId, req.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			response.ErrorResponse(g, response.ErrNotFoundUser)
			return
		}
		response.ErrorResponse(g, response.ErrEmailExists)
		return
	}
	response.SuccessResponse(g, response.AddEmail, nil)
}

// Account godoc
// @Summary      Upgrade to role owner
// @Description  Request here and join with us, create your own eatery
// @Tags         Accounts
// @Accept       json
// @Produce      json
// @Param        request body UpgradeOwnerRequest true "request"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /accounts/upgrade [post]
func (as *AccountController) UpgradeOnwerRequest(g *gin.Context) {
	var req UpgradeOwnerRequest
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, response.ErrBadRequest)
		return
	}
	err := as.service.UpgradeOwnerRequest(g, auth.UserId, req.AccountID)
	if err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"upgrade_queue_account_id_key\" (SQLSTATE 23505)" {
			response.ErrorResponse(g, response.ErrVerify)
			return
		}
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 201, nil)
}
