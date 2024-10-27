package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/jackc/pgx/v5/pgtype"
)

type PostService struct {
	postRepo       *repo.PostRepo
	accountService *AccountService
}

func NewPostRepo(repo *repo.PostRepo, accountService *AccountService) (*PostService, error) {
	return &PostService{
		postRepo:       repo,
		accountService: accountService,
	}, nil
}

func (ps *PostService) CreatePost(ctx context.Context, post_type int32, description string, user_id, account_id int64, x, y string, images []string) (response.PostResponse, error) {
	var res response.PostResponse
	if strings.TrimSpace(description) == "" && len(images) == 0 {
		return res, fmt.Errorf("description and image can't empty")
	}
	var location pgtype.Text
	if x != "" && y != "" {
		location = pgtype.Text{String: fmt.Sprintf("POINT(%s %s)", x, y), Valid: true}
	} else {
		location = pgtype.Text{Valid: false}
	}

	user, err := ps.accountService.GetAccountById(ctx, account_id)
	if err != nil {
		return res, fmt.Errorf("account not found")
	}

	if user_id != user {
		return res, fmt.Errorf("not you")
	}

	post, err := ps.postRepo.CreatePost(ctx, db.CreatePostParams{
		PostTypeID:     post_type,
		AccountID:      account_id,
		Description:    pgtype.Text{String: description, Valid: true},
		StGeomfromtext: location,
	})

	if err != nil {
		return res, err
	}

	var imgs []db.PostImage

	for _, image := range images {
		i, err := ps.postRepo.CreateImagePost(ctx, db.AddImagePostParams{
			UrlImage: image,
			PostID:   post.ID,
		})
		if err != nil {
			return res, err
		}
		imgs = append(imgs, i)
	}
	res = response.PostRes(post, imgs)
	return res, nil
}
