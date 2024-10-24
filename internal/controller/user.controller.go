package controller

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
	config      util.Config
	token       token.Maker
}

func NewUserController(service *service.UserService, config util.Config, token token.Maker) (*UserController, error) {
	return &UserController{
		userService: service,
		config:      config,
		token:       token,
	}, nil
}

func (uc *UserController) Register(g *gin.Context) {
	var req response.RegisterRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 400, 40000)
		return
	}
	res, err := uc.userService.Register(g, req)
	if err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"users_username_key\" (SQLSTATE 23505)" {
			response.ErrorResponse(g, 404, 40900)
			return
		}
		response.ErrorNonKnow(g, 404, err.Error())
		return
	}
	response.SuccessResponse(g, 201, res)
}

func (uc *UserController) Login(g *gin.Context) {
	var req response.LoginRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 400, 40000)
		return
	}
	user, err := uc.userService.Login(g, req.Username, req.Password)
	if err != nil {
		response.ErrorNonKnow(g, 404, err.Error())
		return
	}
	token, err := uc.token.CreateToken(user.ID, user.Username, uc.config.AccessTokenDuration)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	res := response.LoginRes(user, token)
	response.SuccessResponse(g, 200, res)
}
