package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
)

type CommentService struct {
	commentRepo *repo.CommentRepository
	userSevice  *UserService
	postService *PostService
}

func NewCommentService(repo *repo.CommentRepository, userService *UserService, postService *PostService) *CommentService {
	return &CommentService{
		commentRepo: repo,
		userSevice:  userService,
		postService: postService,
	}
}

func (cs *CommentService) CreateComment(ctx context.Context, description string, user_id, post_top_id int64) (response.CommentResponse, error) {
	var res response.CommentResponse
	user, err := cs.userSevice.GetUser(ctx, user_id)
	if err != nil {
		return res, err
	}

	_, err = cs.postService.GetPost(ctx, post_top_id)
	if err != nil {
		return res, err
	}

	comment, err := cs.commentRepo.CreateComment(ctx, description, user_id, post_top_id)

	if err != nil {
		return res, err
	}

	res = response.CommentRes(comment, user)

	return res, nil
}

func (cs *CommentService) ListComment(ctx context.Context, post_id, page, pageSize int64) ([]response.CommentResponse, error) {

	var res []response.CommentResponse
	post, err := cs.postService.GetPost(ctx, post_id)
	if err != nil {
		return res, err
	}
	comments, err := cs.commentRepo.ListComment(ctx, post.ID, int32(page), int32(pageSize))
	if err != nil {
		return res, err
	}

	for _, comment := range comments {
		user, err := cs.userSevice.GetUser(ctx, comment.UserID)
		if err != nil {
			return res, err
		}
		commentRes := response.CommentRes(comment, user)
		res = append(res, commentRes)
	}
	if len(res) == 0 {
		return []response.CommentResponse{}, nil
	}
	return res, nil
}

func (cs *CommentService) UpdateComment(ctx context.Context, id, user_id int64, description string) (response.CommentResponse, error) {
	var res response.CommentResponse
	if strings.TrimSpace(description) == "" {
		return res, fmt.Errorf("description can't empty")
	}
	comment, err := cs.commentRepo.GetCommentById(ctx, id)
	if err != nil {
		return res, err
	}
	if comment.Description.String == description {
		return res, fmt.Errorf("you are not update description")
	}
	if user_id != comment.UserID {
		return res, fmt.Errorf("Forbidden")
	}

	user, err := cs.userSevice.GetUser(ctx, user_id)
	if err != nil {
		return res, err
	}

	update, err := cs.commentRepo.UpdateComment(ctx, id, description)
	if err != nil {
		return res, err
	}

	res = response.CommentRes(update, user)

	return res, nil
}

func (cs *CommentService) DeleteComment(ctx context.Context, id, user_id int64, role_id int32) error {
	comment, err := cs.commentRepo.GetCommentById(ctx, id)
	if err != nil {
		return fmt.Errorf("NotFound")
	}
	if role_id != 1 && comment.UserID != user_id {
		return fmt.Errorf("unauthorize")
	}

	err = cs.commentRepo.DeleteComment(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
