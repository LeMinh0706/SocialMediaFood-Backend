package controller

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/models"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type FollowController struct {
	followService *service.FollowService
}

func NewFollowController(service *service.FollowService) (*FollowController, error) {
	return &FollowController{
		followService: service,
	}, nil
}

func (fc *FollowController) GetFollowStatus(g *gin.Context) {
	var req models.FollowRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	res, err := fc.followService.GetFollowStatus(g, req.FromFollow, req.ToFollow)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 201, res)
}

func (fc *FollowController) GetFollow(g *gin.Context) {
	from_id := g.Query("from_id")
	page := g.Query("page")
	pageSize := g.Query("page_size")
	follow, err := fc.followService.GetFollow(g, from_id, page, pageSize)
	if err != nil {
		GetListErr(g, err)
		return
	}
	response.SuccessResponse(g, 200, follow)
}

func (fc *FollowController) GetFollower(g *gin.Context) {
	from_id := g.Query("from_id")
	page := g.Query("page")
	pageSize := g.Query("page_size")
	follow, err := fc.followService.GetFollower(g, from_id, page, pageSize)
	if err != nil {
		GetListErr(g, err)
		return
	}
	response.SuccessResponse(g, 200, follow)
}

func (fc *FollowController) CreateRequest(g *gin.Context) {
	var req models.FollowRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	res, err := fc.followService.CreateFollowRequest(g, req.FromFollow, req.ToFollow)
	if err != nil {
		GetListErr(g, err)
		return
	}
	response.SuccessResponse(g, 201, res)
}

func (fc *FollowController) GetFriend(g *gin.Context) {
	from_id := g.Query("from_id")
	page := g.Query("page")
	pageSize := g.Query("page_size")
	follow, err := fc.followService.GetFriend(g, from_id, page, pageSize)
	if err != nil {
		GetListErr(g, err)
		return
	}
	response.SuccessResponse(g, 201, follow)
}

func (fc *FollowController) UpdateFriend(g *gin.Context) {
	var req models.FollowRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	res, err := fc.followService.UpdateFriend(g, req.FromFollow, req.ToFollow)
	if err != nil {
		GetListErr(g, err)
		return
	}
	response.SuccessResponse(g, 201, res)
}

func (fc *FollowController) DeleteFollow(g *gin.Context) {
	var req models.FollowRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	err := fc.followService.DeleteFollow(g, req.FromFollow, req.ToFollow)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 204, nil)
}
