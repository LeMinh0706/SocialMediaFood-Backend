package controller

import (
	"strconv"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/models"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
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

func (cc *CommentController) CreateComment(g *gin.Context) {
	var req models.CommentRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40011)
		return
	}
	request := models.CommentReq(req)
	comment, err := cc.commentService.CreateComment(g, request)
	if err != nil {
		if err == pgx.ErrNoRows {
			response.ErrorResponse(g, 40402)
			return
		}
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 201, comment)
}

func (cc *CommentController) GetComment(g *gin.Context) {
	str := g.Param("id")
	id, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 40004)
		return
	}
	comment, err := cc.commentService.GetComment(g, id)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 200, comment)
}

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
