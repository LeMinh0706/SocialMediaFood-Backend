package follower

import (
	"strconv"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/post"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

type FollowerController struct {
	service IFollowerService
	token   token.Maker
}

func NewFollowerController(service IFollowerService, token token.Maker) *FollowerController {
	return &FollowerController{
		service: service,
		token:   token,
	}
}

func (fc *FollowerController) FollowRequest(g *gin.Context) {
	var req CreateFollowRequest
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	follower, err := fc.service.FollowRequest(g, auth.UserId, req.FromID, req.ToID)
	if err != nil {
		FollowErr(g, err)
		return
	}
	response.SuccessResponse(g, 201, follower)
}

func (fc *FollowerController) GetFollowStatus(g *gin.Context) {
	from := g.Query("from_id")
	to := g.Query("to_id")
	from_id, err := strconv.ParseInt(from, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrAccountID)
		return
	}
	to_id, err := strconv.ParseInt(to, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrAccountID)
		return
	}
	status, err := fc.service.GetRequestStatus(g, db.GetFollowStatusParams{FromFollow: from_id, ToFollow: to_id})
	if err != nil {
		FollowErr(g, err)
		return
	}
	response.SuccessResponse(g, 200, status)
}

func (fc *FollowerController) GetFollowType(g *gin.Context) {
	status := g.Query("status")
	from := g.Query("from_id")
	from_id, err := strconv.ParseInt(from, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrAccountID)
		return
	}
	pageStr := g.Query("page")
	pageSizeStr := g.Query("page_size")
	page, pageSize := post.CheckQuery(g, pageStr, pageSizeStr)
	if page == 0 || pageSize == 0 {
		return
	}

	if !StatusCheck(g, status) {
		return
	}

	list, err := fc.service.GetFollowType(g, status, page, pageSize, from_id)
	if err != nil {
		FollowErr(g, err)
		return
	}
	response.SuccessResponse(g, 200, list)
}

func (fc *FollowerController) UpdateFriend(g *gin.Context) {
	var req db.UpdateFriendParams
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	err := fc.service.UpdateStatus(g, auth.UserId, req.FromFollow, req.ToFollow)
	if err != nil {
		FollowErr(g, err)
		return
	}
	response.SuccessResponse(g, 20101, nil)
}

func (fc *FollowerController) UnFollow(g *gin.Context) {
	var req db.DeleteFollowParams
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	if err := fc.service.UnFollow(g, auth.UserId, req.FromFollow, req.ToFollow); err != nil {
		FollowErr(g, err)
		return
	}
	response.SuccessResponse(g, 204, nil)
}
