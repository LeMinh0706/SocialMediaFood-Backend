package service

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
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

func (service *ReactPostService) ListUserReact(ctx context.Context, post_id int64) (response.ReactPostResponse, error) {
	var res response.ReactPostResponse
	var userRes response.UserReactResponse
	_, err := service.postService.GetPost(ctx, post_id)
	if err != nil {
		return res, err
	}
	// var userRes response.UserReactResponse
	users, err := service.reactRepo.GetPostReact(ctx, post_id)
	if err != nil {
		return res, err
	}
	res = response.ReactPostRes(post_id, int64(len(users)))
	for _, user := range users {
		userRes.UserID = user
		res.Users = append(res.Users, userRes)
	}
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
