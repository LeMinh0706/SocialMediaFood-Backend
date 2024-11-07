package controller

import (
	"strconv"

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

// Follower godoc
// @Summary      Get follower status
// @Description  Get status in wall of other person
// @Tags         Follower
// @Accept       json
// @Produce      json
// @Param        from_follow query int true "From Your Account"
// @Param        to_follow query int true "To Account"
// @Security BearerAuth
// @Success      200  {object}  db.Follower
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /follow [get]
func (fc *FollowController) GetFollowStatus(g *gin.Context) {
	FromStr := g.Query("from_follow")
	from_id, err := strconv.ParseInt(FromStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 40004)
		return
	}
	ToStr := g.Query("to_follow")
	to_id, err := strconv.ParseInt(ToStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 40004)
		return
	}
	res, err := fc.followService.GetFollowStatus(g, from_id, to_id)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 201, res)
}

// Follower godoc
// @Summary      Get list follower
// @Description  Get list account you follow
// @Tags         Follower
// @Accept       json
// @Produce      json
// @Param        from_id query int true "Your AccountID"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Security BearerAuth
// @Success      200  {object}  models.ListFollow
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /follow/self [get]
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

// Follower godoc
// @Summary      Get list follower
// @Description  Get list account follow you
// @Tags         Follower
// @Accept       json
// @Produce      json
// @Param        from_id query int true "Your AccountID"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Security BearerAuth
// @Success      200  {object}  models.ListFollow
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /follow/other [get]
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

// Follower godoc
// @Summary      Create follow
// @Description  Create follow for to another user
// @Tags         Follower
// @Accept       json
// @Produce      json
// @Param        request body models.FollowRequest true "request"
// @Security BearerAuth
// @Success      200  {object}  models.FollowRespone
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /follow [post]
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

// Follower godoc
// @Summary      Get list friend
// @Description  Get your friend list
// @Tags         Follower
// @Accept       json
// @Produce      json
// @Param        from_id query int true "Your AccountID"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Security BearerAuth
// @Success      200  {object}  models.ListFollow
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /follow/friend [get]
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

// Follower godoc
// @Summary      Update Friend
// @Description  Make you guy become a friend
// @Tags         Follower
// @Accept       json
// @Produce      json
// @Param        request body models.FollowRequest true "request"
// @Security BearerAuth
// @Success      201  {object} 	models.FollowRespone
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /follow [put]
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

// React godoc
// @Summary      Delete your follow
// @Description  You will unfollow to another and we will delete two record in db
// @Tags         Follower
// @Accept       json
// @Produce      json
// @Param        request body db.DeleteReactParams true "request"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /follow [delete]
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
