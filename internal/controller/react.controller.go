package controller

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type ReactController struct {
	reactService *service.ReactService
}

func NewReactController(service *service.ReactService) (*ReactController, error) {
	return &ReactController{
		reactService: service,
	}, nil
}

func (rc *ReactController) CreateReact(g *gin.Context) {
	var req db.CreateReactParams
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	react, err := rc.reactService.CreateReact(g, req)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 201, react)
}
