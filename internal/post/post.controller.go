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

func (pc *PostController) GetPostById(g *gin.Context) {
	var req struct {
		Id int64 `json:"id" binding:"required, min=1"`
	}
	param := g.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		response.ErrorResponse(g, 400, "Bad request")
		return
	}
	req.Id = id
	post, err := pc.postService.GetPost(g, req.Id)
	if err != nil {
		if err.Error() == "NotFound" {
			response.ErrorResponse(g, 404, "Can not found this post")
			return
		}
		response.ErrorResponse(g, 401, err.Error())
	}
	response.SuccessResponse(g, 200, post)
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
