package controllers

import (
	"fmt"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/services"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService *services.PostService
}

func NewPostController() *PostController {
	return &PostController{
		postService: services.NewPostService(),
	}
}

func (pc *PostController) CreatePost(g *gin.Context) {
	var req struct {
		Description string `json:"description"`
		UserId      int64  `json:"user_id" binding:"required"`
	}

	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 400, fmt.Sprintf("Error: %v", err.Error()))
		return
	}

	post, err := pc.postService.CreatePost(g.Request.Context(), req.Description, req.UserId)
	if err != nil {
		response.ErrorResponse(g, 500, "Failed to create post")
		return
	}
	res := response.PostResponse{ID: post.ID, PostTypeID: post.PostTypeID, UserID: post.UserID, Description: post.Description.String, DateCreatePost: post.DateCreatePost}

	response.SuccessResponse(g, 200, res)
}
