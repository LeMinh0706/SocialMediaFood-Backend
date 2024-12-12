package user

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/handler"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	service IUserService
	config  util.Config
	token   token.Maker
	refesh  token.Maker
}

func NewUserController(service IUserService, config util.Config, token token.Maker, refesh token.Maker) *UserController {
	return &UserController{
		service: service,
		config:  config,
		token:   token,
		refesh:  refesh,
	}
}

// User godoc
// @Summary      Login user
// @Description  Login to be more handsome
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body LoginRequest true "request"
// @Success      200  {object}  LoginResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /users/login [post]
func (uc *UserController) Login(g *gin.Context) {
	var req LoginRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		handler.ValidateRegister(g, err)
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
	tokenId, _ := uuid.NewRandom()

	token, err := uc.token.CreateToken(tokenId, user.Username, uc.config.AccessTokenDuration)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	refreshId, _ := uuid.NewRandom()
	refesh, err := uc.refesh.CreateToken(refreshId, user.Username, uc.config.RefeshTokenDuration)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
	}
	res := LoginResponse{AccessToken: token, RefeshToken: refesh}
	response.SuccessResponse(g, 200, res)
}

// User godoc
// @Summary      Register user
// @Description  Join with us
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body db.RegisterRequest true "request"
// @Success      200  {object}  RegisterResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /users/register [post]
func (uc *UserController) RegisterTx(g *gin.Context) {
	var req db.RegisterRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		handler.ValidateRegister(g, err)
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
		if err.Error() == "this mail is invalid" {
			response.ErrorResponse(g, response.ErrEmailInvalid)
			return
		}
		response.ErrorNonKnow(g, 401, err.Error())
		return
	}
	response.SuccessResponse(g, 201, res)
}

// User godoc
// @Summary      Refresh token for user
// @Description  Join with us
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body AccessRequest true "request"
// @Success      200  {object}  RegisterResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /users/refresh [post]
func (uc *UserController) RefeshToken(g *gin.Context) {
	var req AccessRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	payload, err := uc.refesh.VerifyToken(req.RefreshToken)
	if err != nil {
		response.ErrorResponse(g, response.ErrTokenInvalid)
		return
	}
	tokenId, _ := uuid.NewRandom()

	token, err := uc.token.CreateToken(tokenId, payload.Username, uc.config.AccessTokenDuration)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 201, AccessResponse{AccessToken: token})
}
