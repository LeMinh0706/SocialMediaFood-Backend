package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/models"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/repo"
	"github.com/jackc/pgx/v5/pgtype"
)

type PostService struct {
	postRepo       *repo.PostRepo
	accountService *AccountService
}

func NewPostService(repo *repo.PostRepo, accountService *AccountService) (*PostService, error) {
	return &PostService{
		postRepo:       repo,
		accountService: accountService,
	}, nil
}

func (ps *PostService) CreatePost(ctx context.Context, post_type int32, description string, user_id, account_id int64, x, y string, images []string) (models.PostResponse, error) {
	var res models.PostResponse
	if strings.TrimSpace(description) == "" && len(images) == 0 {
		return res, fmt.Errorf("description and image can't empty")
	}
	var location pgtype.Text
	if x != "" && y != "" {
		location = pgtype.Text{String: fmt.Sprintf("POINT(%s %s)", x, y), Valid: true}
	} else {
		location = pgtype.Text{Valid: false}
	}

	acc, err := ps.accountService.GetAccountForAction(ctx, user_id, account_id)
	if err != nil {
		return res, err
	}

	if user_id != acc.UserID {
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
	res = models.PostRes(post, acc, imgs, 0)
	return res, nil
}

func (ps *PostService) GetPost(ctx context.Context, id int64) (models.PostResponse, error) {
	var res models.PostResponse

	post, err := ps.postRepo.GetPost(ctx, id)
	if err != nil {
		return res, err
	}
	acc, err := ps.accountService.GetAccountById(ctx, post.AccountID)
	if err != nil {
		return res, err
	}
	img, err := ps.postRepo.GetImagePost(ctx, post.ID)
	if err != nil {
		return res, err
	}
	like, err := ps.postRepo.CountLike(ctx, id)
	if err != nil {
		return res, err
	}
	res = models.PostRes(db.CreatePostRow(post), acc, img, like)

	return res, nil
}

func (ps *PostService) GetListPost(ctx context.Context, pageStr, pageSizeStr string) ([]models.PostResponse, error) {
	var res []models.PostResponse
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		return []models.PostResponse{}, fmt.Errorf("page number")
	}
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		return []models.PostResponse{}, fmt.Errorf("pagesize number")
	}

	post_id, err := ps.postRepo.GetListPost(ctx, int32(page), int32(pageSize))
	if err != nil {
		return []models.PostResponse{}, err
	}

	for _, id := range post_id {
		post, err := ps.GetPost(ctx, id)
		if err != nil {
			return []models.PostResponse{}, err
		}
		res = append(res, post)
	}

	if len(res) == 0 {
		return []models.PostResponse{}, nil
	}

	return res, nil
}

func (ps *PostService) DeletePost(ctx context.Context, id, user_id int64) error {
	post, err := ps.GetPost(ctx, id)
	if err != nil {
		return err
	}

	_, err = ps.accountService.GetAccountForAction(ctx, user_id, post.AccountID)
	if err != nil {
		return err
	}

	err = ps.postRepo.DeletePost(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PostService) DeleteImage(ctx context.Context, id int64) error {
	_, err := ps.postRepo.GetImageById(ctx, id)
	if err != nil {
		return err
	}
	err = ps.postRepo.DeleteImage(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PostService) UpdatePost(ctx context.Context, description string, user_id, id int64) (models.PostResponse, error) {
	var res models.PostResponse
	post, err := ps.GetPost(ctx, id)
	if err != nil {
		return res, err
	}
	acc, err := ps.accountService.GetAccountForAction(ctx, user_id, post.AccountID)
	if err != nil {
		return res, err
	}
	update, err := ps.postRepo.UpdatePost(ctx, db.UpdatePostParams{ID: id, Description: pgtype.Text{String: description, Valid: true}})
	if err != nil {
		return res, err
	}
	res = models.UpdatePostRes(update, acc, post.Images)
	return res, nil
}

func (ps *PostService) GetUserPost(ctx context.Context, pageStr, pageSizeStr, account_idStr string) ([]models.PostResponse, error) {
	var res []models.PostResponse

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		return []models.PostResponse{}, fmt.Errorf("page number")
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		return []models.PostResponse{}, fmt.Errorf("pagesize number")
	}

	account_id, err := strconv.ParseInt(account_idStr, 10, 64)
	if err != nil {
		return []models.PostResponse{}, fmt.Errorf("account_id number")
	}

	_, err = ps.accountService.GetAccountById(ctx, account_id)
	if err != nil {
		return []models.PostResponse{}, err
	}

	list, err := ps.postRepo.GetUserPost(ctx, int32(page), int32(pageSize), account_id)
	if err != nil {
		return []models.PostResponse{}, err
	}
	for _, id := range list {
		post, err := ps.GetPost(ctx, id)
		if err != nil {
			return []models.PostResponse{}, err
		}
		res = append(res, post)
	}
	if len(res) == 0 {
		return []models.PostResponse{}, nil

	}
	return res, nil
}
