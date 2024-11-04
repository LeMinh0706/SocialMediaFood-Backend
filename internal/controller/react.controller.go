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

func (rc *ReactController) GetReactPost(g *gin.Context) {
	page := g.Query("page")
	pageSize := g.Query("page_size")
	post_id := g.Query("post_id")

	reacts, err := rc.reactService.GetReactPost(g, post_id, page, pageSize)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 200, reacts)
}

func (rc *ReactController) UpdateReact(g *gin.Context) {
	var req db.UpdateStateParams
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	update, err := rc.reactService.UpdateState(g, req.AccountID, req.PostID, req.State)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 201, update)
}

func (rc *ReactController) UnlikePost(g *gin.Context) {
	var req db.DeleteReactParams
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	err := rc.reactService.UnlikePost(g, req.AccountID, req.PostID)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 204, nil)
}
