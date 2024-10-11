package service

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
)

type ReactPostService struct {
	reactRepo   *repo.ReactPostRepository
	userService *UserService
	postService *PostService
}

func NewReactPostService(reactRepo *repo.ReactPostRepository, userService *UserService, postService *PostService) *ReactPostService {
	return &ReactPostService{
		reactRepo:   reactRepo,
		userService: userService,
		postService: postService,
	}
}

func (service *ReactPostService) ReactPost(ctx context.Context, arg db.CreateReactParams) (db.ReactPost, error) {
	var res db.ReactPost
	_, err := service.userService.GetUser(ctx, arg.UserID)
	if err != nil {
		return res, err
	}
	_, err = service.postService.GetPost(ctx, arg.PostID)
	if err != nil {
		return res, err
	}
	react, err := service.reactRepo.CreateReact(ctx, arg)
	if err != nil {
		return res, err
	}
	res = react
	return res, nil
}

func (service *ReactPostService) UnLikePost(ctx context.Context, arg db.GetReactParams) error {
	react, err := service.reactRepo.GetReact(ctx, arg)
	if err != nil {
		return err
	}
	err = service.reactRepo.DeleteReact(ctx, react.ID)
	if err != nil {
		return err
	}
	return nil
}
