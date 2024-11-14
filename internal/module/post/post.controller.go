package post

import (
	"strconv"
	"strings"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type PostController struct {
	service IPostService
	token   token.Maker
}

func NewPostController(service IPostService, token token.Maker) *PostController {
	return &PostController{
		service: service,
		token:   token,
	}
}

// Post godoc
// @Summary      Create post
// @Description  Create post
// @Tags         Posts
// @Accept       multipart/form-data
// @Produce      json
// @Param        description formData string false "Description"
// @Param        account_id formData string true "Account ID"
// @Param        lng formData string false "Lng"
// @Param        lat formData string false "Lat"
// @Param        images formData []file false "Images post"
// @Security BearerAuth
// @Success      200  {object}  PostResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /posts [post]
func (pc *PostController) CreatePost(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	accStr := g.PostForm("account_id")
	account_id, err := strconv.ParseInt(accStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 40004)
		return
	}
	description := g.PostForm("description")
	lng := g.PostForm("lng")
	lat := g.PostForm("lat")

	if !CheckValidPosition(g, lng, lat) {
		return
	}
	form, err := g.MultipartForm()
	if err != nil {
		response.ErrorResponse(g, 40000)
		return
	}

	files := form.File["images"]

	images, code := AddImageFileError(g, 4, files)
	if code > 40000 {
		response.ErrorResponse(g, code)
		return
	}

	if strings.TrimSpace(description) == "" && len(images) == 0 {
		response.ErrorResponse(g, 40022)
		return
	}

	post, err := pc.service.CreatePost(g, description, lat, lng, images, account_id, auth.UserId)
	if err != nil {
		if err.Error() == "not you" {
			response.ErrorResponse(g, response.ErrYourSelf)
			return
		}
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}

	response.SuccessResponse(g, 201, post)
}

// Post godoc
// @Summary      Get list post
// @Description  Get list post with page and page size (Limit-Offset)
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Success      200  {object}  []PostResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /posts [get]
func (pc *PostController) GetListPost(g *gin.Context) {
	pageStr := g.Query("page")
	pageSizeStr := g.Query("page_size")
	page, pageSize := CheckQuery(g, pageStr, pageSizeStr)
	if page == 0 || pageSize == 0 {
		return
	}
	list, err := pc.service.GetListPost(g, page, pageSize)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 200, list)
}

// Post godoc
// @Summary      Get list post
// @Description  Get list post with account_id, page and page size (Limit-Offset)
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Param        account_id query int true "Account ID"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Security BearerAuth
// @Success      200  {object}  []PostResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /posts/person [get]
func (pc *PostController) GetPersonPost(g *gin.Context) {
	accStr := g.Query("account_id")
	account_id, err := strconv.ParseInt(accStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 40004)
		return
	}
	pageStr := g.Query("page")
	pageSizeStr := g.Query("page_size")
	page, pageSize := CheckQuery(g, pageStr, pageSizeStr)
	if page == 0 || pageSize == 0 {
		return
	}
	list, err := pc.service.GetPersonPost(g, account_id, page, pageSize)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 200, list)
}

// Post godoc
// @Summary      Update post
// @Description  Just update content post
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Param        description formData string false "Description"
// @Security BearerAuth
// @Success      201  {object} 	PostResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /posts/{id} [put]
func (pc *PostController) UpdateContentPost(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	idStr := g.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrBadRequestId)
		return
	}
	description := g.PostForm("description")
	update, err := pc.service.UpdateContentPost(g, description, id, auth.UserId)
	if err != nil {
		CheckPostStringError(g, err)
		return
	}
	response.SuccessResponse(g, 201, update)
}

// Post godoc
// @Summary      Delete Image
// @Description  Delete image when update post (maybe)
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /posts/images/{id} [delete]
func (pc *PostController) DeleteImage(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	idStr := g.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrBadRequestId)
		return
	}
	err = pc.service.DeleteImage(g, id, auth.UserId)
	if err != nil {
		CheckPostStringError(g, err)
		return
	}
	response.SuccessResponse(g, 204, nil)
}

// Post godoc
// @Summary      Delete post
// @Description  Just Delete post
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /posts/soft-delete/{id} [post]
func (pc *PostController) DeletePost(g *gin.Context) {
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	idStr := g.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrBadRequestId)
		return
	}
	err = pc.service.DeletePost(g, id, auth.UserId)
	if err != nil {
		CheckPostStringError(g, err)
		return
	}
	response.SuccessResponse(g, 204, nil)
}

func CheckPostStringError(g *gin.Context, err error) {
	if err == pgx.ErrNoRows {
		response.ErrorResponse(g, response.ErrFindPost)
		return
	}
	if err.Error() == "not you" {
		response.ErrorResponse(g, response.ErrYourSelf)
		return
	}
	if err.Error() == "ERROR: duplicate key value violates unique constraint \"react_post_post_id_account_id_idx\" (SQLSTATE 23505)" {
		response.ErrorResponse(g, response.ErrLike)
		return
	}
	if err.Error() == "err like" {
		response.ErrorResponse(g, response.ErrUnlike)
		return
	}
	response.ErrorNonKnow(g, 500, err.Error())
}

func (pc *PostController) GetHomePagePost(g *gin.Context) {
	accStr := g.Query("account_id")
	account_id, err := strconv.ParseInt(accStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 40004)
		return
	}
	pageStr := g.Query("page")
	pageSizeStr := g.Query("page_size")
	page, pageSize := CheckQuery(g, pageStr, pageSizeStr)
	if page == 0 || pageSize == 0 {
		return
	}
	list, err := pc.service.GetHomePagePost(g, account_id, page, pageSize)
	if err != nil {
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 200, list)
}
