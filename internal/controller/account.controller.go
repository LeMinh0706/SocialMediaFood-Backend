package controller

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accountService *service.AccountService
	token          token.Maker
}

func NewAccountController(service *service.AccountService, token token.Maker) (*AccountController, error) {
	return &AccountController{
		accountService: service,
		token:          token,
	}, nil
}

func (ac *AccountController) GetMe(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	me, err := ac.accountService.GetAccountUser(g, auth.UserId)
	if err != nil {
		response.ErrorNonKnow(g, 400, err.Error())
		return
	}
	response.SuccessResponse(g, 200, me)
}
