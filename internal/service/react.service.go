package service

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
)

type ReactService struct {
	reactRepo *repo.ReactRepo
}

func NewReactService(repo *repo.ReactRepo) (*ReactService, error) {
	return &ReactService{
		reactRepo: repo,
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
