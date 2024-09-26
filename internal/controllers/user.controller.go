package controllers

import (
	"fmt"
	"strconv"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/services"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) Register(g *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 400, fmt.Sprintf("Can not signup, error: %v", err))
		return
	}

	user, err := uc.userService.Register(g.Request.Context(), req.Username, req.Password)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_username_key\"" {
			response.ErrorResponse(g, 401, "User exist")
			return
		}
		response.ErrorResponse(g, 500, fmt.Sprint(err.Error()))
		return
	}

	res := response.RegisterResponse{ID: user.ID, Email: user.Email.String, Fullname: user.Fullname, Username: user.Username, Gender: user.Gender, RoleID: user.RoleID, DateCreateAccount: user.DateCreateAccount}
	response.SuccessResponse(g, 201, res)
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

	res := response.UserResponse{ID: user.ID, Fullname: user.Fullname, Gender: user.Gender, RoleID: user.RoleID, DateCreateAccount: user.DateCreateAccount}
	response.SuccessResponse(g, 200, res)
}
