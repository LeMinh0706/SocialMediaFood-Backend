package controller

import (
	"database/sql"

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
