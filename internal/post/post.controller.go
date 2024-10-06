package post

import (
	"fmt"
	"strconv"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService *PostService
	token       token.Maker
}

func NewPostController(tokenMaker token.Maker) *PostController {
	return &PostController{
		postService: NewPostService(),
		token:       tokenMaker,
	}
}

func (pc *PostController) CreatePost(g *gin.Context) {
	// var req response.PostRequest

	authPayload := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)

	description := g.PostForm("description")

	form, err := g.MultipartForm()
	if err != nil {
		response.ErrorResponse(g, 400, err.Error())
		return
	}
	var images []string

	files := form.File["images"]

	const maxSize = 6 << 20

	for _, file := range files {
		if file.Size > maxSize {
			response.ErrorResponse(g, 401, "File too large, only allowed 8MB")
			return
		}
		filename := fmt.Sprintf("upload/post/%d_%s", time.Now().Unix(), file.Filename)
		if !middlewares.FileUploadCheck(filename) {
			response.ErrorResponse(g, 400, "Can only use file .png, .jpg, .jpeg, .gif")
			return
		}

		if err := g.SaveUploadedFile(file, filename); err != nil {
			response.ErrorResponse(g, 500, err.Error())
			return
		}
		images = append(images, filename)
	}
	post, err := pc.postService.CreatePost(g, description, authPayload.UserId, images)
	if err != nil {
		response.ErrorResponse(g, 401, err.Error())
		return
	}

	response.SuccessResponse(g, 201, post)
}

func (pc *PostController) GetListPost(g *gin.Context) {
	pageStr := g.Query("page")
	pageSizeStr := g.Query("page_size")
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 400, "Bad request, page should be number")
		return
	}
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 400, "Bad request, page_size should be number")
		return
	}
	posts, err := pc.postService.GetListPost(g, page, pageSize)
	if err != nil {
		response.ErrorResponse(g, 401, err.Error())
		return
	}

	response.SuccessResponse(g, 200, posts)
}
