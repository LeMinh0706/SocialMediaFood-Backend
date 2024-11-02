package controller

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/models"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/gin-gonic/gin"
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
		response.ErrorNonKnow(g, 400, err.Error())
		return
	}
	response.SuccessResponse(g, 201, comment)
}
