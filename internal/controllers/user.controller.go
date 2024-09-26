package controllers

import (
	"fmt"

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
		response.ErrorResponse(g, 500, "Failed to signup")
		return
	}

	res := response.UserResponse{ID: user.ID, Email: user.Email.String, Fullname: user.Fullname, Username: user.Username, Gender: user.Gender, RoleID: user.RoleID, DateCreateAccount: user.DateCreateAccount}
	response.SuccessResponse(g, 201, res)
}
