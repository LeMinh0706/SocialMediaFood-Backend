package post

import (
	"strconv"
	"strings"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	service IPostService
	token   token.Maker
}

func NewPostController(service IPostService, token token.Maker) (*PostController, error) {
	return &PostController{
		service: service,
		token:   token,
	}, nil
}

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
		response.ErrorNonKnow(g, 500, err.Error())
		return
	}

	response.SuccessResponse(g, 201, post)
}

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
