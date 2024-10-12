package controller

import (
	"database/sql"
	"strconv"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService *service.CommentService
	token          token.Maker
}

func NewCommentController(tokenMaker token.Maker, commentService *service.CommentService) *CommentController {
	return &CommentController{
		commentService: commentService,
		token:          tokenMaker,
	}
}

// User godoc
// @Summary      Create comment
// @Description  Create comment in post
// @Tags         comment
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param        request body response.CommentRequest true "request"
// @Success      201  {object}  response.CommentResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /comment [post]
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

// Comment godoc
// @Summary      Get list comment
// @Description  Get list comment with page and page size (Limit-Offset)
// @Tags         comment
// @Accept       json
// @Produce      json
// @Param        post_id query int true "Post Id"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page size"
// @Success      200  {object}  []response.CommentResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /comment [get]
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

// Comment godoc
// @Summary      Update comment
// @Description  Update comment
// @Tags         comment
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param        id path int true "Post Id"
// @Param		 request body response.UpdateCommentRequest true "request"
// @Success      200  {object}  response.CommentResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /comment/{id} [put]
func (cc *CommentController) UpdateComment(g *gin.Context) {
	var req response.UpdateCommentRequest
	authPayload := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)

	param := g.Param("id")

	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 400, 40004)
		return
	}

	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 400, 40000)
		return
	}

	comment, err := cc.commentService.UpdateComment(g, id, authPayload.UserId, req.Description)
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

// Comment godoc
// @Summary      Delete comment
// @Description  Delete comment with id
// @Tags         comment
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param        id path int true "Post Id"
// @Success      204  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /comment/{id} [delete]
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
