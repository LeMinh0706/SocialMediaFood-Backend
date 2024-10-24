package controller

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(service *service.UserService) (*UserController, error) {
	return &UserController{
		userService: service,
	}, nil
}

func (uc *UserController) Register(g *gin.Context) {
	var req db.RegisterParams
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 400, 40000)
		return
	}
	res, err := uc.userService.Register(g, req.Username, req.HashPashword, req.Email)
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
