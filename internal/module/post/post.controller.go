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
	if err != nil {
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
		return
	}
}
