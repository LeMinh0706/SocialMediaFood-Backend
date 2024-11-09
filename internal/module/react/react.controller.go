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
