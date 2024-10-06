package comment

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService *CommentService
	token          token.Maker
}

func NewCommentController(tokenMaker token.Maker) *CommentController {
	return &CommentController{
		commentService: NewCommentService(),
		token:          tokenMaker,
	}
}

func (cc *CommentController) CreateComment(g *gin.Context) {
	var req response.CommentRequest

	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 400, err.Error())
		return
	}

	res, err := cc.commentService.CreateComment(g, req.Description, req.UserID, req.PostTopID)
	if err != nil {
		response.ErrorResponse(g, 401, err.Error())
		return
	}

	response.SuccessResponse(g, 201, res)
}
