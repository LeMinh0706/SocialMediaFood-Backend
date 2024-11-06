package controller

import (
	"strconv"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/models"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type PostController struct {
	postService *service.PostService
}

func NewPostController(service *service.PostService) (*PostController, error) {
	return &PostController{
		postService: service,
	}, nil
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
// @Success      200  {object}  models.PostResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /posts [post]
func (pc *PostController) CreatePost(g *gin.Context) {

	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	account_id := g.PostForm("account_id")
	id, err := strconv.ParseInt(account_id, 10, 64)
	if err != nil {
		response.ErrorNonKnow(g, 400, err.Error())
		return
	}
	x := g.PostForm("lng")
	y := g.PostForm("lat")
	if (x == "" && y != "") || (x != "" && y == "") {
		response.ErrorResponse(g, 40013)
		return
	}

	description := g.PostForm("description")

	form, err := g.MultipartForm()
	if err != nil {
		response.ErrorResponse(g, 40000)
		return
	}

	files := form.File["images"]

	images, err := AddImageFileError(g, 4, files)
	if err != nil {
		response.ErrorNonKnow(g, 400, err.Error())
		return
	}

	post, err := pc.postService.CreatePost(g, 1, description, auth.UserId, id, x, y, images)
	if err != nil {
		if err.Error() == "not you" {
			response.ErrorResponse(g, 40103)
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
// @Param        page_size query int true "Page size"
// @Success      200  {object}  []models.PostResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /posts [get]
func (pc *PostController) GetListPost(g *gin.Context) {
	page := g.Query("page")
	pageSize := g.Query("page_size")

	posts, err := pc.postService.GetListPost(g, page, pageSize)
	if err != nil {
		GetListErr(g, err)
		return
	}
	response.SuccessResponse(g, 200, posts)
}

// Post godoc
// @Summary      Get list post
// @Description  Get list post with account_id, page and page size (Limit-Offset)
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Param        account_id query int true "Account ID"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page size"
// @Success      200  {object}  []models.PostResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /posts/person [get]
func (pc *PostController) GetUserPost(g *gin.Context) {
	account_id := g.Query("account_id")
	page := g.Query("page")
	pageSize := g.Query("page_size")
	posts, err := pc.postService.GetUserPost(g, page, pageSize, account_id)
	if err != nil {
		GetListErr(g, err)
		return
	}
	response.SuccessResponse(g, 200, posts)
}

func (pc *PostController) GetPost(g *gin.Context) {
	str := g.Param("id")
	id, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 40004)
		return
	}
	post, err := pc.postService.GetPost(g, id)
	if err != nil {
		response.ErrorResponse(g, 40402)
		return
	}
	response.SuccessResponse(g, 200, post)
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

	str := g.Param("id")
	id, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 40004)
		return
	}
	if err = pc.postService.DeletePost(g, id, auth.UserId); err != nil {
		if err.Error() == "not you" {
			response.ErrorResponse(g, 40103)
			return
		}
		response.ErrorResponse(g, 40402)
		return
	}

	response.SuccessResponse(g, 204, nil)
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
func (pc *PostController) DeleteImagePost(g *gin.Context) {
	str := g.Param("id")
	id, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 40004)
		return
	}
	if err = pc.postService.DeleteImage(g, id); err != nil {
		if err == pgx.ErrNoRows {
			response.ErrorResponse(g, 40408)
			return
		}
		response.ErrorResponse(g, 40402)
		return
	}
	response.SuccessResponse(g, 204, nil)
}

// Post godoc
// @Summary      Update post
// @Description  Just update content post
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Param        request body models.UpdatePostRequest true "request"
// @Security BearerAuth
// @Success      201  {object} 	models.PostResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /posts/ [put]
func (pc *PostController) UpdatePost(g *gin.Context) {
	var req models.UpdatePostRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)

	update, err := pc.postService.UpdatePost(g, req.Description, auth.UserId, req.ID)
	if err != nil {
		if err.Error() == "not you" {
			response.ErrorResponse(g, 40103)
			return
		}
		if err == pgx.ErrNoRows {
			response.ErrorResponse(g, 40402)
			return
		}
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 200, update)
}
