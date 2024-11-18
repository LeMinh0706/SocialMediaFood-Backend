package comment

import (
	"strconv"
	"strings"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/post"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	service ICommentService
	token   token.Maker
}

func NewCommentController(service ICommentService, token token.Maker) *CommentController {
	return &CommentController{
		service: service,
		token:   token,
	}
}

// Comment godoc
// @Summary      Create comment
// @Description  Create comment
// @Tags         Comments
// @Accept       multipart/form-data
// @Produce      json
// @Param        post_id formData string true "PostID"
// @Param        description formData string true "Description"
// @Param        account_id formData string true "AccountID"
// @Param        image formData file false "Image comment"
// @Security BearerAuth
// @Success      201  {object}  CommentResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /comments [post]
func (cc *CommentController) CreateComment(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	description := g.PostForm("description")
	accountStr := g.PostForm("account_id")
	post_idStr := g.PostForm("post_id")
	if strings.TrimSpace(description) == "" {
		response.ErrorResponse(g, response.ContentNull)
		return
	}
	account_id, err := strconv.ParseInt(accountStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrAccountID)
		return
	}

	post_id, err := strconv.ParseInt(post_idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrBadRequestId)
		return
	}
	var file string
	image, err := g.FormFile("image")
	if err == nil {
		var code int
		file, code = SaveCommentImage(g, image)
		if code >= 40000 {
			response.ErrorResponse(g, code)
			return
		}
	}
	comment, err := cc.service.CreateComment(g, account_id, auth.UserId, post_id, description, file)
	if err != nil {
		post.CheckPostStringError(g, err)
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
// @Param        post_id query int true "PostID"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Security BearerAuth
// @Success      200  {object}  []CommentResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /comments [get]
func (cc *CommentController) GetListComment(g *gin.Context) {
	postIdStr := g.Query("post_id")
	pageStr := g.Query("page")
	pageSizeStr := g.Query("page_size")
	page, pageSize := post.CheckQuery(g, pageStr, pageSizeStr)
	if page == 0 || pageSize == 0 {
		return
	}
	post_top_id, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrBadRequestId)
		return
	}
	comments, err := cc.service.GetListComment(g, page, pageSize, post_top_id)
	if err != nil {
		post.CheckPostStringError(g, err)
		return
	}
	response.SuccessResponse(g, 200, comments)
}

// Comment godoc
// @Summary      Create comment
// @Description  Create comment
// @Tags         Comments
// @Accept       multipart/form-data
// @Produce      json
// @Param        id formData string true "CommentID"
// @Param        description formData string true "Description"
// @Param        image formData file false "Image comment"
// @Security BearerAuth
// @Success      201  {object}  CommentResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /comments [put]
func (cc *CommentController) UpdateComment(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	idStr := g.PostForm("id")
	description := g.PostForm("description")
	if strings.TrimSpace(description) == "" {
		response.ErrorResponse(g, response.ContentNull)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrBadRequestId)
		return
	}

	var file string
	image, err := g.FormFile("image")
	if err == nil {
		var code int
		file, code = SaveCommentImage(g, image)
		if code >= 40000 {
			response.ErrorResponse(g, code)
			return
		}
	}
	comment, err := cc.service.UpdateComment(g, auth.UserId, id, description, file)
	if err != nil {
		post.CheckPostStringError(g, err)
		return
	}
	response.SuccessResponse(g, 201, comment)
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
func (cc *CommentController) DeleteComment(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	idStr := g.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrBadRequestId)
		return
	}
	err = cc.service.DeleteComment(g, id, auth.UserId)
	if err != nil {
		post.CheckPostStringError(g, err)
		// response.ErrorResponse(g, 40119)
		return
	}
	response.SuccessResponse(g, 204, nil)
}
