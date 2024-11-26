package react

import (
	"strconv"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/handler"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
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
// @Success      201  {object}  db.ReactPost
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /react [post]
func (rc *ReactController) CreateReact(g *gin.Context) {
	var req ReactRequest
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	if req.State == 0 {
		req.State = 1
	}
	react, err := rc.service.CreateReact(g, auth.UserId, req.AccountID, req.PostID, req.State)
	if err != nil {
		handler.CheckPostStringError(g, err)
		return
	}
	response.SuccessResponse(g, 201, react)
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
	page, pageSize := handler.CheckQuery(g, pageStr, pageSizeStr)
	if page == 0 || pageSize == 0 {
		return
	}
	reacts, err := rc.service.GetListReactPost(g, page, pageSize, id)
	if err != nil {
		handler.CheckPostStringError(g, err)
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
		handler.CheckPostStringError(g, err)
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
		handler.CheckPostStringError(g, err)
		return
	}
	response.SuccessResponse(g, 204, nil)
}
