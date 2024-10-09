package comment

import (
	"database/sql"
	"strconv"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
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

	authPayload := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)

	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorNonKnow(g, 400, err.Error())
		return
	}

	if authPayload.UserId != req.UserID {
		response.ErrorResponse(g, 403, 40103)
		return
	}

	res, err := cc.commentService.CreateComment(g, req.Description, req.UserID, req.PostTopID)
	if err != nil {
		if err.Error() == "NotFound" {
			response.ErrorResponse(g, 404, 40402)
			return
		}
		response.ErrorNonKnow(g, 404, err.Error())
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
		response.ErrorResponse(g, 400, 40004)
		return
	}
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 400, 40001)
		return
	}
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 400, 40002)
		return
	}

	comments, err := cc.commentService.ListComment(g, postId, page, pageSize)
	if err != nil {
		if err.Error() == "NotFound" {
			response.ErrorResponse(g, 404, 40402)
			return
		}
		response.ErrorNonKnow(g, 404, err.Error())
		return
	}
	response.SuccessResponse(g, 200, comments)
}

func (cc *CommentController) UpdateComment(g *gin.Context) {
	var req response.UpdateCommentRequest
	authPayload := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)

	param := g.Param("id")

	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 400, 40004)
		return
	}

	req.ID = id

	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 400, 40000)
		return
	}

	comment, err := cc.commentService.UpdateComment(g, req.ID, authPayload.UserId, req.Description)
	if err != nil {

		if err == sql.ErrNoRows {
			response.ErrorResponse(g, 404, 40402)
			return
		}
		if err.Error() == "Forbidden" {
			response.ErrorResponse(g, 403, 40103)
			return
		}
		response.ErrorNonKnow(g, 400, err.Error())
		return
	}
	response.SuccessResponse(g, 201, comment)
}

func (cc *CommentController) DeleteComment(g *gin.Context) {
	param := g.Param("id")
	authPayload := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)

	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 400, 40000)
		return
	}
	err = cc.commentService.DeleteComment(g, id, authPayload.UserId, authPayload.RoleID)
	if err != nil {
		switch err.Error() {
		case "unauthorize":
			response.ErrorResponse(g, 401, 40103)
		case "NotFound":
			response.ErrorResponse(g, 404, 40402)
		case "sql: no rows in result set":
			response.ErrorResponse(g, 404, 40402)
		default:
			response.ErrorNonKnow(g, 500, err.Error())
		}
		return
	}

	response.SuccessResponse(g, 204, nil)
}
