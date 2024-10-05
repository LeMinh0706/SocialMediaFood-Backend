package post

import (
	"fmt"
	"strconv"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService *PostService
}

func NewPostController() *PostController {
	return &PostController{
		postService: NewPostService(),
	}
}

func (pc *PostController) CreatePost(g *gin.Context) {
	// var req response.PostRequest

	description := g.PostForm("description")
	userId := g.PostForm("user_id")
	uid, err := strconv.ParseInt(userId, 10, 64)

	if err != nil {
		response.ErrorResponse(g, 400, fmt.Sprintf("Error: %v", err.Error()))
		return
	}

	form, err := g.MultipartForm()
	if err != nil {
		response.ErrorResponse(g, 400, fmt.Sprintf("Error %v", err.Error()))
		return
	}

	var images []string
	if form != nil {
		files := form.File["images"]
		for _, file := range files {
			filename := fmt.Sprintf("upload/post/%d_%s", time.Now().Unix(), file.Filename)
			if err := g.SaveUploadedFile(file, filename); err != nil {
				response.ErrorResponse(g, 500, fmt.Sprintf("Error, %v", err.Error()))
				return
			}
			images = append(images, filename)
		}
	}

	if description == "" && len(images) == 0 {
		response.ErrorResponse(g, 400, "description or images can not empty")
		return
	}

	post, err := pc.postService.CreatePost(g.Request.Context(), description, uid, images)
	if err != nil {
		response.ErrorResponse(g, 500, err.Error())
		return
	}

	response.SuccessResponse(g, 200, post)
}
