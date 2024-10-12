package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService *service.PostService
	token       token.Maker
}

func NewPostController(tokenMaker token.Maker, postService *service.PostService) *PostController {
	return &PostController{
		postService: postService,
		token:       tokenMaker,
	}
}

// Post godoc
// @Summary      Create post
// @Description  Create post
// @Tags         Posts
// @Accept       multipart/form-data
// @Produce      json
// @Param        description formData string false "Description"
// @Param        images formData []file false "Images post"
// @Security BearerAuth
// @Success      200  {object}  response.PostResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /post [post]
func (pc *PostController) CreatePost(g *gin.Context) {
	// var req response.PostRequest

	authPayload := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)

	description := g.PostForm("description")

	form, err := g.MultipartForm()
	if err != nil {
		response.ErrorNonKnow(g, 400, err.Error())
		return
	}
	var images []string

	files := form.File["images"]

	const maxSize = 6 << 20

	for _, file := range files {
		if file.Size > maxSize {
			response.ErrorResponse(g, 413, 41300)
			return
		}
		filename := fmt.Sprintf("upload/post/%d_%s", time.Now().Unix(), file.Filename)
		if !util.FileUploadCheck(filename) {
			response.ErrorResponse(g, 400, 40003)
			return
		}

		if err := g.SaveUploadedFile(file, filename); err != nil {
			response.ErrorNonKnow(g, 500, err.Error())
			return
		}
		images = append(images, filename)
	}
	post, err := pc.postService.CreatePost(g, description, authPayload.UserId, images)
	if err != nil {
		response.ErrorNonKnow(g, 404, err.Error())
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
// @Success      200  {object}  []response.PostResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /post [get]
func (pc *PostController) GetListPost(g *gin.Context) {
	pageStr := g.Query("page")
	pageSizeStr := g.Query("page_size")
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
	posts, err := pc.postService.GetListPost(g, page, pageSize)
	if err != nil {
		response.ErrorNonKnow(g, 404, err.Error())
		return
	}

	response.SuccessResponse(g, 200, posts)
}
