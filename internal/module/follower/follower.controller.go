package follower

import (
	"strconv"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/handler"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
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

// Follower godoc
// @Summary      Create follow
// @Description  Create follow for to another user
// @Tags         Follower
// @Accept       json
// @Produce      json
// @Param        request body CreateFollowRequest true "request"
// @Security BearerAuth
// @Success      201  {object}  FollowResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /follower [post]
func (fc *FollowerController) FollowRequest(g *gin.Context) {
	var req CreateFollowRequest
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	follower, err := fc.service.FollowRequest(g, auth.UserId, req.FromID, req.ToID)
	if err != nil {
		handler.FollowErr(g, err)
		return
	}
	response.SuccessResponse(g, 201, follower)
}

// Follower godoc
// @Summary      Get status from your and the other
// @Description  To see the relationship from your to another
// @Tags         Follower
// @Accept       json
// @Produce      json
// @Param        from_id query int true "From your"
// @Param        to_id query int true "To person"
// @Security BearerAuth
// @Success      200  {object}  db.GetFollowStatusRow
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /follower/status [get]
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
		handler.FollowErr(g, err)
		return
	}
	response.SuccessResponse(g, 200, status)
}

// Follower godoc
// @Summary      Get list Follower
// @Description  Get list follower from user with 3 types: "accept", "request", "friend"
// @Tags         Follower
// @Accept       json
// @Produce      json
// @Param        status query string true "Status"
// @Param        from_id query int true "From Account"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Security BearerAuth
// @Success      200  {object}  ListFollow
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /follower [get]
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
	page, pageSize := handler.CheckQuery(g, pageStr, pageSizeStr)
	if page == 0 || pageSize == 0 {
		return
	}

	if !handler.StatusCheck(g, status) {
		return
	}

	list, err := fc.service.GetFollowType(g, status, page, pageSize, from_id)
	if err != nil {
		handler.FollowErr(g, err)
		return
	}
	response.SuccessResponse(g, 200, list)
}

// Follower godoc
// @Summary      Update Friend
// @Description  Make you guy become a friend
// @Tags         Follower
// @Accept       json
// @Produce      json
// @Param        request body db.UpdateFriendParams true "request"
// @Security BearerAuth
// @Success      201  "no content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /follower [put]
func (fc *FollowerController) UpdateFriend(g *gin.Context) {
	var req db.UpdateFriendParams
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	err := fc.service.UpdateStatus(g, auth.UserId, req.FromFollow, req.ToFollow)
	if err != nil {
		handler.FollowErr(g, err)
		return
	}
	response.SuccessResponse(g, 20101, nil)
}

// Follower godoc
// @Summary      Delete your follow
// @Description  You will unfollow to another and we will delete two record in db
// @Tags         Follower
// @Accept       json
// @Produce      json
// @Param        request body db.DeleteFollowParams true "request"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /follower [delete]
func (fc *FollowerController) UnFollow(g *gin.Context) {
	var req db.DeleteFollowParams
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	if err := fc.service.UnFollow(g, auth.UserId, req.FromFollow, req.ToFollow); err != nil {
		handler.FollowErr(g, err)
		return
	}
	response.SuccessResponse(g, 204, nil)
}
