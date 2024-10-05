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

	userId := g.PostForm("user_id")
	description := g.PostForm("description")
	uid, err := strconv.ParseInt(userId, 10, 64)

	if err != nil {
		response.ErrorResponse(g, 400, fmt.Sprintf("Error: %v", err.Error()))
		return
	}

	if uid != authPayload.UserId {
		response.ErrorResponse(g, 400, "It's not your, you cant create post for others")
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

	post, err := pc.postService.CreatePost(g.Request.Context(), description, uid, images)
	if err != nil {
		response.ErrorResponse(g, 401, err.Error())
		return
	}

	response.SuccessResponse(g, 200, post)
}
