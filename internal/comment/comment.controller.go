package comment

import (
	"strconv"

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

func (cc *CommentController) ListComment(g *gin.Context) {
	postIdStr := g.Query("post_id")
	pageStr := g.Query("page")
	pageSizeStr := g.Query("page_size")
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 404, "Bad request")
		return
	}
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 404, "Bad request")
		return
	}
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 404, "Bad request")
		return
	}

	comments, err := cc.commentService.ListComment(g, postId, page, pageSize)
	if err != nil {
		response.ErrorResponse(g, 401, err.Error())
		return
	}
	response.SuccessResponse(g, 200, comments)
}
