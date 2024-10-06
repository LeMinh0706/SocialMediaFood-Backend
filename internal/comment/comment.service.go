package comment

import (
	"context"
	"fmt"
	"log"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/post"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/user"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
)

type CommentService struct {
	commentRepo *CommentRepository
}

func NewCommentService() *CommentService {
	repo, err := NewCommentRepo()
	if err != nil {
		log.Fatal("Error:", err)
	}
	return &CommentService{
		commentRepo: repo,
	}
}

func (cs *CommentService) CreateComment(ctx context.Context, description string, user_id, post_top_id int64) (response.CommentResponse, error) {

	user, err := user.NewUserService().GetUser(ctx, user_id)
	if err != nil {
		return response.CommentResponse{}, err
	}

	post, err := post.NewPostService().GetPost(ctx, post_top_id)
	if err != nil {
		return response.CommentResponse{}, err
	}

	if post.PostTypeID == 2 {
		return response.CommentResponse{}, fmt.Errorf("Can not create in comment")
	}

	comment, err := cs.commentRepo.CreateComment(ctx, description, user_id, post_top_id)

	if err != nil {
		return response.CommentResponse{}, err
	}

	res := response.CommentRes(comment, user)

	return res, nil
}
