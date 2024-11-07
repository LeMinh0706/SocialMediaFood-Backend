package controller

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
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

// React godoc
// @Summary      Create reacttion
// @Description  Create reaction for post
// @Tags         React
// @Accept       json
// @Produce      json
// @Param        request body db.CreateReactParams true "request"
// @Security BearerAuth
// @Success      200  {object}  db.ReactPost
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /react [post]
func (rc *ReactController) CreateReact(g *gin.Context) {
	var req db.CreateReactParams
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	react, err := rc.reactService.CreateReact(g, req, auth.UserId)
	if err != nil {
		GetListErr(g, err)
		return
	}
	response.SuccessResponse(g, 201, react)
}

// Post godoc
// @Summary      Get list reactions
// @Description  Get list reactions with post_id, page and page size (Limit-Offset)
// @Tags         React
// @Accept       json
// @Produce      json
// @Param        post_id query int true "PostID"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Security BearerAuth
// @Success      200  {object}  models.ListReactResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /react [get]
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

// Post godoc
// @Summary      Update react
// @Description  Just update reaction type 1 for like, 2 for hearth, 3 for sad, 4 for angry
// @Tags         React
// @Accept       json
// @Produce      json
// @Param        request body db.UpdateStateParams true "request"
// @Security BearerAuth
// @Success      201  {object} 	db.ReactPost
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /react [put]
func (rc *ReactController) UpdateReact(g *gin.Context) {
	var req db.UpdateStateParams
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	update, err := rc.reactService.UpdateState(g, req.AccountID, req.PostID, req.State, auth.UserId)
	if err != nil {
		GetListErr(g, err)
		return
	}
	response.SuccessResponse(g, 201, update)
}

// Post godoc
// @Summary      Delete Reactions
// @Description  Delete your reaction with post_id
// @Tags         React
// @Accept       json
// @Produce      json
// @Param        request body db.DeleteReactParams true "request"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /react [delete]
func (rc *ReactController) UnlikePost(g *gin.Context) {
	var req db.DeleteReactParams
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	err := rc.reactService.UnlikePost(g, req.AccountID, req.PostID, auth.UserId)
	if err != nil {
		GetListErr(g, err)
		return
	}
	response.SuccessResponse(g, 204, nil)
}
