package account

import (
	"strconv"

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
// @Param        image formData file true "Image comment"
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
		file, code = SaveAccountImage(g, image)
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
