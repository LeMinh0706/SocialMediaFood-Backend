package controller

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
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
		response.ErrorResponse(g, 400, 40000)
		return
	}

	files := form.File["images"]
	if len(files) > 4 {
		response.ErrorResponse(g, 400, 40005)
		return
	}
	images, err := AddImageFileError(g, files)
	if err != nil {
		response.ErrorNonKnow(g, 400, err.Error())
		return
	}

	post, err := pc.postService.CreatePost(g, 1, description, auth.UserId, id, x, y, images)
	if err != nil {
		if err.Error() == "not you" {
			response.ErrorResponse(g, 401, 40103)
			return
		}
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}
	response.SuccessResponse(g, 201, post)
}

func AddImageFileError(g *gin.Context, files []*multipart.FileHeader) ([]string, error) {
	const maxSize = 4 << 20
	for _, file := range files {
		if !util.FileExtCheck(file.Filename) {
			return nil, fmt.Errorf("only accept .jpeg/.jpg/.png/.gif")
		}
	}

	var images []string
	for i, file := range files {
		if file.Size > maxSize {
			return nil, fmt.Errorf("images size must less than 4 Mb")
		}
		filename := fmt.Sprintf("upload/post/%d_%d%s", time.Now().Unix(), i, filepath.Ext(file.Filename))
		if err := g.SaveUploadedFile(file, filename); err != nil {
			return nil, err
		}
		images = append(images, filename)
	}
	return images, nil
}
