package react

import (
	"strconv"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/post"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

type ReactController struct {
	service IReactService
	token   token.Maker
}

func NewReactController(service IReactService, token token.Maker) *ReactController {
	return &ReactController{
		service: service,
		token:   token,
	}
}

// React godoc
// @Summary      Create reacttion
// @Description  Create reaction for post
// @Tags         React
// @Accept       json
// @Produce      json
// @Param        request body ReactRequest true "request"
// @Security BearerAuth
// @Success      200  {object}  db.ReactPost
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /react [post]
func (rc *ReactController) CreateReact(g *gin.Context) {
	var req ReactRequest
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	react, err := rc.service.CreateReact(g, auth.UserId, req.AccountID, req.PostID)
	if err != nil {
		post.CheckPostStringError(g, err)
		return
	}
	response.SuccessResponse(g, 201, react)
}

// React godoc
// @Summary      Get reactions
// @Description  Get with post_id, to say have you liked this post
// @Tags         React
// @Accept       json
// @Produce      json
// @Param        account_id query int true "AccountID"
// @Param        post_id query int true "PostID"
// @Security BearerAuth
// @Success      200  {object}  ReactResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /react [get]
func (rc *ReactController) GetReactPost(g *gin.Context) {
	accStr := g.Query("account_id")
	postStr := g.Query("post_id")
	account_id, err := strconv.ParseInt(accStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrAccountID)
		return
	}
	post_id, err := strconv.ParseInt(postStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrBadRequestId)
		return
	}
	react, err := rc.service.GetReactPost(g, account_id, post_id)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 200, react)
}

// React godoc
// @Summary      Get list reactions
// @Description  Get list reactions with post_id, page and page size (Limit-Offset)
// @Tags         React
// @Accept       json
// @Produce      json
// @Param        id path int true "PostID"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Security BearerAuth
// @Success      200  {object}  ListReactResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /react/post/{id} [get]
func (rc *ReactController) GetListReact(g *gin.Context) {
	idStr := g.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrBadRequestId)
		return
	}
	pageStr := g.Query("page")
	pageSizeStr := g.Query("page_size")
	page, pageSize := post.CheckQuery(g, pageStr, pageSizeStr)
	if page == 0 || pageSize == 0 {
		return
	}
	reacts, err := rc.service.GetListReactPost(g, page, pageSize, id)
	if err != nil {
		post.CheckPostStringError(g, err)
		return
	}

	response.SuccessResponse(g, 200, reacts)
}

// React godoc
// @Summary      Change reaction state
// @Description  Just change reaction type 1 for like, 2 for hearth, 3 for sad, 4 for angry
// @Tags         React
// @Accept       json
// @Produce      json
// @Param        request body db.UpdateStateParams true "request"
// @Security BearerAuth
// @Success      201  {object} 	db.ReactPost
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /react [put]
func (rc *ReactController) ChangeReactState(g *gin.Context) {
	var req db.UpdateStateParams
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	update, err := rc.service.ChangeReactState(g, auth.UserId, req.AccountID, req.PostID, req.State)
	if err != nil {
		post.CheckPostStringError(g, err)
		return
	}
	response.SuccessResponse(g, 201, update)
}

// React godoc
// @Summary      Unlike post
// @Description  Delete your reaction with on any post
// @Tags         React
// @Accept       json
// @Produce      json
// @Param        request body db.DeleteReactParams true "request"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /react [delete]
func (rc *ReactController) UnReaction(g *gin.Context) {
	var req db.DeleteReactParams
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	err := rc.service.UnReaction(g, auth.UserId, req.AccountID, req.PostID)
	if err != nil {
		post.CheckPostStringError(g, err)
		return
	}
	response.SuccessResponse(g, 204, nil)
}
