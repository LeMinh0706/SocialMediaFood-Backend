package controller

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/models"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type CommentController struct {
	commentService *service.CommentService
}

func NewCommentController(service *service.CommentService) (*CommentController, error) {
	return &CommentController{
		commentService: service,
	}, nil
}

// Comment godoc
// @Summary      Create comment
// @Description  Create comment
// @Tags         Comments
// @Accept       json
// @Produce      json
// @Param        request body models.CommentRequest true "request"
// @Security BearerAuth
// @Success      200  {object}  models.CommentResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /comments [post]
func (cc *CommentController) CreateComment(g *gin.Context) {
	var req models.CommentRequest
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40011)
		return
	}
	request := models.CommentReq(req)
	comment, err := cc.commentService.CreateComment(g, auth.UserId, request)
	if err != nil {
		if err == pgx.ErrNoRows {
			response.ErrorResponse(g, 40402)
			return
		}
		GetListErr(g, err)
		// response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 201, comment)
}

// Post godoc
// @Summary      Get list post
// @Description  Get list post with post_top_id, page and page size (Limit-Offset)
// @Tags         Comments
// @Accept       json
// @Produce      json
// @Param        post_id query int true "Post ID"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page size"
// @Success      200  {object}  []models.CommentResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /comments [get]
func (cc *CommentController) GetListComment(g *gin.Context) {
	post_id := g.Query("post_id")
	page := g.Query("page")
	pageSize := g.Query("page_size")
	comments, err := cc.commentService.GetListComment(g, page, pageSize, post_id)
	if err != nil {
		GetListErr(g, err)
		return
	}
	response.SuccessResponse(g, 200, comments)
}

// Post godoc
// @Summary      Update Comment
// @Description  Just update content post
// @Tags         Comments
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Param        request body models.UpdateCommentRequest true "request"
// @Security BearerAuth
// @Success      201  {object} 	models.CommentResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /comments/{id} [put]
func (cc *CommentController) UpdateComment(g *gin.Context) {
	var req models.UpdateCommentRequest
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40011)
		return
	}
	id := g.Param("id")
	update, err := cc.commentService.UpdateComment(g, id, auth.UserId, req)
	if err != nil {
		GetListErr(g, err)
		return
	}
	response.SuccessResponse(g, 201, update)
}

// Post godoc
// @Summary      Delete Comment
// @Description  Delete comment
// @Tags         Comments
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /comments/{id} [delete]
func (cc *CommentController) DeteleComment(g *gin.Context) {
	id := g.Param("id")
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	err := cc.commentService.DeleteComment(g, id, auth.UserId)
	if err != nil {
		GetListErr(g, err)
		return
	}
	response.SuccessResponse(g, 204, nil)
}

// func (cc *CommentController) GetComment(g *gin.Context) {
// 	str := g.Param("id")
// 	id, err := strconv.ParseInt(str, 10, 64)
// 	if err != nil {
// 		response.ErrorResponse(g, 40004)
// 		return
// 	}
// 	comment, err := cc.commentService.GetComment(g, id)
// 	if err != nil {
// 		response.ErrorNonKnow(g, 500, err.Error())
// 		return
// 	}
// 	response.SuccessResponse(g, 200, comment)
// }
