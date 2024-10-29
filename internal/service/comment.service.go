package service

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
)

type CommentService struct {
	commentRepo *repo.CommentRepo
}

func NewCommentService(repo *repo.CommentRepo) (*CommentService, error) {
	return &CommentService{
		commentRepo: repo,
	}, nil
}

func (cs *CommentService) CreateComment(ctx context.Context, arg db.CreateCommentParams) (db.CreateCommentRow, error) {
	var res db.CreateCommentRow
	comment, err := cs.commentRepo.CreateComment(ctx, arg)
	if err != nil {
		return res, err
	}
	return comment, nil
}
