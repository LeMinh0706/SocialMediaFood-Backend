package controller

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
	tokenMaker  token.Maker
	config      util.Config
}

func NewUserController(tokenMaker token.Maker, userSerive *service.UserService, config util.Config) *UserController {
	return &UserController{
		userService: userSerive,
		tokenMaker:  tokenMaker,
		config:      config,
	}
}

// User godoc
// @Summary      Register user
// @Description  Join with us
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        request body response.RequestResponse true "request"
// @Success      200  {object}  response.RegisterResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /accounts/register [post]
func (uc *UserController) Register(g *gin.Context) {
	var req response.RequestResponse
	if err := g.ShouldBindJSON(&req); err != nil {
		if err.Error() == "Key: 'RequestResponse.Password' Error:Field validation for 'Password' failed on the 'min' tag" {
			response.ErrorNonKnow(g, 400, "Password too short, must be more than 6 character")
			return
		} else if err.Error() == "Key: 'RequestResponse.Password' Error:Field validation for 'Password' failed on the 'max' tag" {
			response.ErrorNonKnow(g, 400, "Password too short, must be less than 18 characters")
			return
		}
		response.ErrorNonKnow(g, 400, err.Error())
		return
	}

	user, err := uc.userService.Register(g.Request.Context(), req.Username, req.Password, req.Fullname, req.Email, req.Gender)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_username_key\"" {
			response.ErrorResponse(g, 409, 40900)
			return
		}
		response.ErrorNonKnow(g, 500, fmt.Sprint(err.Error()))
		return
	}

	res := response.RegisterRes(user)
	response.SuccessResponse(g, 201, res)
}

// User godoc
// @Summary      Login user
// @Description  Login to be more handsome
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        request body response.RequestLogin true "request"
// @Success      200  {object}  response.LoginResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /accounts/login [post]
func (uc *UserController) Login(g *gin.Context) {
	var req response.RequestLogin

	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 400, 40000)
		return
	}

	user, err := uc.userService.Login(g, req.Username, req.Password)

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			response.ErrorResponse(g, 404, 40400)
			return
		} else if err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
			response.ErrorResponse(g, 404, 40400)
			return
		}
		response.ErrorNonKnow(g, 404, err.Error())
		return
	}

	token, err := uc.tokenMaker.CreateToken(user.ID, user.RoleID, req.Username, uc.config.AccessTokenDuration)

	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}

	res := response.LoginResponse{AccessToken: token, User: user}

	response.SuccessResponse(g, 200, res)
}

// User godoc
// @Summary      It's you
// @Description  Login to be more handsome
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Success      200  {object}  response.UserResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /accounts/me [get]
func (uc *UserController) GetMe(g *gin.Context) {

	authPayload := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)

	me, err := uc.userService.GetMe(g.Request.Context(), authPayload.Username)
	if err != nil {
		response.ErrorResponse(g, 404, 40401)
	}

	response.SuccessResponse(g, 200, me)
}

func (uc *UserController) GetById(g *gin.Context) {
	var req struct {
		Id int64 `json:"id" binding:"required"`
	}
	param := g.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 400, 40000)
		return
	}

	req.Id = id

	user, err := uc.userService.GetUser(g.Request.Context(), req.Id)
	if err != nil {
		response.ErrorResponse(g, 404, 40401)
		return
	}

	response.SuccessResponse(g, 200, user)
}
