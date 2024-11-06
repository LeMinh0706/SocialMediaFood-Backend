package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/models"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
)

type ReactService struct {
	reactRepo      *repo.ReactRepo
	accountService *AccountService
}

func NewReactService(repo *repo.ReactRepo, accountService *AccountService) (*ReactService, error) {
	return &ReactService{
		reactRepo:      repo,
		accountService: accountService,
	}, nil
}

func (rs *ReactService) CreateReact(ctx context.Context, arg db.CreateReactParams) (db.ReactPost, error) {
	var res db.ReactPost
	react, err := rs.reactRepo.CreateReact(ctx, db.CreateReactParams{
		AccountID: arg.AccountID,
		PostID:    arg.PostID,
		State:     1,
	})
	if err != nil {
		return res, err
	}
	return react, nil
}

func (rs *ReactService) GetReactPost(ctx context.Context, post_idStr, pageStr, pageSizeStr string) (models.ListReactResponse, error) {
	var list []models.ReactResponse
	var res models.ListReactResponse
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		return res, fmt.Errorf("page number")
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		return res, fmt.Errorf("pagesize number")
	}

	post_id, err := strconv.ParseInt(post_idStr, 10, 64)
	if err != nil {
		return res, fmt.Errorf("pagesize number")
	}

	reacts, err := rs.reactRepo.GetReactPost(ctx, int32(page), int32(pageSize), post_id)
	if err != nil {
		return res, err
	}

	for _, react := range reacts {
		acc, err := rs.accountService.GetAccountById(ctx, react.AccountID)
		if err != nil {
			return res, err
		}
		result := models.ReactResponse{ID: react.ID, Account: acc}
		list = append(list, result)
	}
	total, err := rs.reactRepo.CountLike(ctx, post_id)
	if err != nil {
		return res, err
	}
	res = models.ListReactResponse{React: list, Total: total}
	return res, nil
}

func (rs *ReactService) GetReact(ctx context.Context, post_id, account_id int64) (int64, error) {
	id, err := rs.reactRepo.GetReact(ctx, db.GetReactParams{AccountID: account_id, PostID: post_id})
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (rs *ReactService) UpdateState(ctx context.Context, account_id, post_id int64, state int32) (db.ReactPost, error) {
	update, err := rs.reactRepo.UpdateState(ctx, post_id, account_id, state)
	if err != nil {
		return db.ReactPost{}, err
	}
	return update, nil
}

func (rs *ReactService) UnlikePost(ctx context.Context, account_id, post_id int64) error {
	err := rs.reactRepo.UnlikePost(ctx, post_id, account_id)
	if err != nil {
		return err
	}
	return nil
}
