package user

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service IUserService
	config  util.Config
	token   token.Maker
}

func NewUserController(service IUserService, config util.Config, token token.Maker) *UserController {
	return &UserController{
		service: service,
		config:  config,
		token:   token,
	}
}

func (uc *UserController) Login(g *gin.Context) {
	var req LoginRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		ValidateRegister(g, err)
		// response.ErrorNonKnow(g, 400, err.Error())
		return
	}
	user, err := uc.service.Login(g, req.Username, req.Password)
	if err != nil {
		if err.Error() == response.WrongUsername {
			response.ErrorResponse(g, 40104)
			return
		}
		if err.Error() == response.WrongPassword {
			response.ErrorResponse(g, 40105)
			return
		}
		response.ErrorNonKnow(g, 404, err.Error())
		return
	}
	token, err := uc.token.CreateToken(user.ID, user.Username, uc.config.AccessTokenDuration)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	res := LoginResponse{Token: token}
	response.SuccessResponse(g, 200, res)
}

func (uc *UserController) RegisterTx(g *gin.Context) {
	var req db.RegisterRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		ValidateRegister(g, err)
		return
	}
	res, err := uc.service.Register(g, req)
	if err != nil {
		if err.Error() == response.UserExists {
			response.ErrorResponse(g, 40900)
			return
		}
		if err.Error() == response.EmailExists {
			response.ErrorResponse(g, 40901)
			return
		}
		response.ErrorNonKnow(g, 401, err.Error())
		return
	}
	response.SuccessResponse(g, 201, res)

}