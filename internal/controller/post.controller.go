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

func (pc *PostController) CreatePost(g *gin.Context) {

	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	account_id := g.PostForm("account_id")
	id, err := strconv.ParseInt(account_id, 10, 64)
	if err != nil {
		response.ErrorNonKnow(g, 400, err.Error())
		return
	}
	x := g.PostForm("direct_x")
	y := g.PostForm("direct_y")

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

func (pc *PostController) GetListPost(g *gin.Context) {
	page := g.Query("page")
	pageSize := g.Query("page_size")

	posts, err := pc.postService.GetListPost(g, page, pageSize)
	if err != nil {
		GetListErr(g, err)
	}
	response.SuccessResponse(g, 200, posts)
}

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

func (pc *PostController) DeleteImagePost(g *gin.Context) {
	str := g.Param("id")
	id, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 40004)
		return
	}
	if err = pc.postService.DeleteImage(g, id); err != nil {
		response.ErrorResponse(g, 40402)
		return
	}
	response.SuccessResponse(g, 204, nil)
}

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
