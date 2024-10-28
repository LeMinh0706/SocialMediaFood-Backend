package controller

import (
	"strconv"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/service"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
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
