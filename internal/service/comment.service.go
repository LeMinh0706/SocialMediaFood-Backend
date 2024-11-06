package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/models"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
	"github.com/jackc/pgx/v5/pgtype"
)

type CommentService struct {
	commentRepo    *repo.CommentRepo
	accountService *AccountService
	postService    *PostService
}

func NewCommentService(repo *repo.CommentRepo, accountService *AccountService, postService *PostService) (*CommentService, error) {
	return &CommentService{
		commentRepo:    repo,
		accountService: accountService,
		postService:    postService,
	}, nil
}

func (cs *CommentService) CreateComment(ctx context.Context, user_id int64, arg db.CreateCommentParams) (models.CommentResponse, error) {
	var res models.CommentResponse

	_, err := cs.postService.GetPost(ctx, arg.PostTopID.Int64)
	if err != nil {
		return res, err
	}
	acc, err := cs.accountService.GetAccountForAction(ctx, user_id, arg.AccountID)
	if err != nil {
		return res, err
	}

	comment, err := cs.commentRepo.CreateComment(ctx, arg)
	if err != nil {
		return res, err
	}

	res = models.CommentRes(acc, db.GetCommentRow(comment))
	return res, nil
}

func (cs *CommentService) GetComment(ctx context.Context, id int64) (models.CommentResponse, error) {
	var res models.CommentResponse
	comment, err := cs.commentRepo.GetComment(ctx, id)
	if err != nil {
		return res, nil
	}

	acc, err := cs.accountService.GetAccountById(ctx, comment.AccountID)
	if err != nil {
		return res, nil
	}

	res = models.CommentRes(acc, comment)
	return res, nil
}

func (cs *CommentService) GetListComment(ctx context.Context, pageStr, pageSizeStr, idStr string) ([]models.CommentResponse, error) {
	var res []models.CommentResponse
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		return []models.CommentResponse{}, fmt.Errorf("page number")
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		return []models.CommentResponse{}, fmt.Errorf("pagesize number")
	}

	post_id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return []models.CommentResponse{}, fmt.Errorf("post_id number")
	}

	_, err = cs.postService.GetPost(ctx, post_id)
	if err != nil {
		return nil, err
	}

	list, err := cs.commentRepo.GetListComment(ctx, int32(page), int32(pageSize), post_id)
	if err != nil {
		return []models.CommentResponse{}, err
	}
	for _, id := range list {
		comment, err := cs.GetComment(ctx, id)
		if err != nil {
			return []models.CommentResponse{}, err
		}
		res = append(res, comment)
	}
	if len(res) == 0 {
		return []models.CommentResponse{}, nil
	}
	return res, nil
}

func (cs *CommentService) UpdateComment(ctx context.Context, idStr string, user_id int64, req models.UpdateCommentRequest) (models.CommentResponse, error) {
	var res models.CommentResponse
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return res, fmt.Errorf("id number")
	}
	comment, err := cs.GetComment(ctx, id)
	if err != nil {
		return res, err
	}

	acc, err := cs.accountService.GetAccountForAction(ctx, user_id, comment.Account.ID)
	if err != nil {
		return res, err
	}

	update, err := cs.commentRepo.UpdateComment(ctx, db.UpdateCommentParams{
		ID:          id,
		Description: pgtype.Text{String: req.Description, Valid: true},
	})
	if err != nil {
		return res, err
	}
	res = models.CommentRes(acc, db.GetCommentRow(update))
	return res, nil
}

func (cs *CommentService) DeleteComment(ctx context.Context, idStr string, user_id int64) error {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fmt.Errorf("id number")
	}
	comment, err := cs.GetComment(ctx, id)
	if err != nil {
		return err
	}
	if comment.Account.UserID != user_id {
		return fmt.Errorf("not you")
	}
	err = cs.commentRepo.DeleteComment(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
