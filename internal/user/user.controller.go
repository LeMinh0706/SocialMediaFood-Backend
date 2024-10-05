package user

import (
	"fmt"
	"strconv"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *UserService
	tokenMaker  token.Maker
}

func NewUserController(tokenMaker token.Maker) *UserController {
	return &UserController{
		userService: NewUserService(),
		tokenMaker:  tokenMaker,
	}
}

func (uc *UserController) Register(g *gin.Context) {
	var req response.RequestResponse
	if err := g.ShouldBindJSON(&req); err != nil {
		if err.Error() == "Key: 'RequestResponse.Password' Error:Field validation for 'Password' failed on the 'min' tag" {
			response.ErrorResponse(g, 400, "Password too short, must be more than 6 character")
			return
		} else if err.Error() == "Key: 'RequestResponse.Password' Error:Field validation for 'Password' failed on the 'max' tag" {
			response.ErrorResponse(g, 400, "Password too short, must be less than 18 characters")
			return
		}
		response.ErrorResponse(g, 400, "Bad request")
		return
	}

	user, err := uc.userService.Register(g.Request.Context(), req.Username, req.Password, req.Gender)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_username_key\"" {
			response.ErrorResponse(g, 409, "User already exist")
			return
		}
		response.ErrorResponse(g, 500, fmt.Sprint(err.Error()))
		return
	}

	res := response.RegisterResponse{ID: user.ID, Email: user.Email.String, Fullname: user.Fullname, Username: user.Username, Gender: user.Gender, RoleID: user.RoleID, DateCreateAccount: user.DateCreateAccount}
	response.SuccessResponse(g, 201, res)
}

func (uc *UserController) Login(g *gin.Context) {
	var req response.RequestLogin

	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 400, err.Error())
		return
	}

	user, err := uc.userService.Login(g, req.Username, req.Password)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			response.ErrorResponse(g, 404, "Wrong username")
			return
		} else if err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
			response.ErrorResponse(g, 404, "Wrong password")
			return
		}
		response.ErrorResponse(g, 404, err.Error())
		return
	}

	token, err := uc.tokenMaker.CreateToken(user.ID, user.RoleID, req.Username, 5*time.Minute)

	if err != nil {
		response.ErrorResponse(g, 500, err.Error())
		return
	}

	res := response.LoginResponse{AccessToken: token, User: user}

	response.SuccessResponse(g, 200, res)
}

func (uc *UserController) GetMe(g *gin.Context) {

	authPayload := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)

	me, err := uc.userService.GetMe(g.Request.Context(), authPayload.Username)
	if err != nil {
		response.ErrorResponse(g, 404, "Cant not found user!")
	}

	response.SuccessResponse(g, 200, me)
}

func (uc *UserController) GetById(g *gin.Context) {
	var req struct {
		Id int64 `json:"id" binding:"required"`
	}
	idParam := g.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 400, fmt.Sprintf("Bad request: %v", err))
		return
	}

	req.Id = id

	user, err := uc.userService.GetUser(g.Request.Context(), req.Id)
	if err != nil {
		response.ErrorResponse(g, 404, "Cant not found user!")
	}

	response.SuccessResponse(g, 200, user)
}