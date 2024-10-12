package controller

import (
	"database/sql"
	"strconv"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

type ReactController struct {
	service *service.ReactPostService
	token   token.Maker
}

func NewReactPostController(token token.Maker, service *service.ReactPostService) *ReactController {
	return &ReactController{
		service: service,
		token:   token,
	}
}

// User godoc
// @Summary      React post
// @Description  React post from user to post
// @Tags         react
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param        request body db.CreateReactParams true "request"
// @Success      201  {object}  db.ReactPost
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /react [post]
func (rc *ReactController) LikePost(g *gin.Context) {
	var req db.CreateReactParams

	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 400, 40000)
		return
	}
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if req.UserID != auth.UserId {
		response.ErrorResponse(g, 400, 40103)
		return
	}
	res, err := rc.service.ReactPost(g, req)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"react_post_post_id_user_id_idx\"" {
			response.ErrorResponse(g, 404, 40404)
			return
		}
		response.ErrorNonKnow(g, 404, err.Error())
		return
	}
	response.SuccessResponse(g, 201, res)
}

// Comment godoc
// @Summary      Delete react
// @Description  Delete react where user id and post id exist, also it's exist in react model
// @Tags         react
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param		 request body db.GetReactParams true "request"
// @Success      204  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /react [delete]
func (rc *ReactController) UnlikePost(g *gin.Context) {
	var req db.GetReactParams

	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 400, 40000)
		return
	}

	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if req.UserID != auth.UserId {
		response.ErrorResponse(g, 400, 40103)
		return
	}

	err := rc.service.UnLikePost(g, req)
	if err != nil {
		if err == sql.ErrNoRows {
			response.ErrorResponse(g, 404, 40403)
			return
		}
	}
	response.SuccessResponse(g, 204, nil)
}

// Comment godoc
// @Summary      Get list react
// @Description  Get list react with no limit offset
// @Tags         react
// @Accept       json
// @Produce      json
// @Param        id path int true "Post Id"
// @Success      200  {object}  response.ReactPostResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /react/{id} [get]
func (rc *ReactController) ListReactPost(g *gin.Context) {
	param := g.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 400, 40000)
		return
	}
	res, err := rc.service.ListUserReact(g, id)
	if err != nil {
		response.ErrorNonKnow(g, 404, err.Error())
		return
	}
	response.SuccessResponse(g, 200, res)
}
