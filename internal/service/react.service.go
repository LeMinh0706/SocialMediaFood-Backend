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

func (rs *ReactService) GetReactPost(ctx context.Context, post_idStr, pageStr, pageSizeStr string) ([]models.ReactResponse, error) {
	var res []models.ReactResponse
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		return []models.ReactResponse{}, fmt.Errorf("page number")
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		return []models.ReactResponse{}, fmt.Errorf("pagesize number")
	}

	post_id, err := strconv.ParseInt(post_idStr, 10, 64)
	if err != nil {
		return []models.ReactResponse{}, fmt.Errorf("pagesize number")
	}

	reacts, err := rs.reactRepo.GetReactPost(ctx, int32(page), int32(pageSize), post_id)
	if err != nil {
		return []models.ReactResponse{}, err
	}

	for _, react := range reacts {
		acc, err := rs.accountService.GetAccountById(ctx, react.AccountID)
		if err != nil {
			return []models.ReactResponse{}, err
		}
		accRes := models.AccountPost(acc)
		result := models.ReactResponse{ID: react.ID, Account: accRes}
		res = append(res, result)
	}

	return res, nil
}
