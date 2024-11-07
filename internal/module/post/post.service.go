package post

import (
	"context"
	"fmt"
	"sync"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/account"
	"github.com/jackc/pgx/v5/pgtype"
)

type PostService struct {
	queries        *db.Queries
	accountService account.IAccountService
}

// GetListPost implements IPostService.
func (p *PostService) GetListPost(ctx context.Context, page int32, pageSize int32) ([]PostResponse, error) {
	var res []PostResponse
	list, err := p.queries.GetListPost(ctx, db.GetListPostParams{Limit: pageSize, Offset: (page - 1) * pageSize})
	if err != nil {
		return res, err
	}
	for _, element := range list {
		post, err := p.GetPost(ctx, element)
		if err != nil {
			return []PostResponse{}, err
		}
		res = append(res, post)
	}
	return res, nil
}

// GetPersonPost implements IPostService.
func (p *PostService) GetPersonPost(ctx context.Context, acoount_id int64, page int32, pageSize int32) ([]PostResponse, error) {
	var res []PostResponse
	list, err := p.queries.GetPersonPost(ctx, db.GetPersonPostParams{AccountID: acoount_id, Limit: pageSize, Offset: (page - 1) * pageSize})
	if err != nil {
		return res, err
	}
	for _, element := range list {
		post, err := p.GetPost(ctx, element)
		if err != nil {
			return []PostResponse{}, err
		}
		res = append(res, post)
	}
	return res, nil
}

// Backup implements IPostService.
func (p *PostService) Backup(ctx context.Context) {
	panic("unimplemented")
}

// CreatePost implements IPostService.
func (p *PostService) CreatePost(ctx context.Context, description string, lng string, lat string, images []string, account_id int64, user_id int64) (PostResponse, error) {
	var res PostResponse
	var wg sync.WaitGroup
	var location pgtype.Text
	var acc db.GetAccountByIdRow
	var err error
	wg.Add(1)
	go func() {
		defer wg.Done()
		if lat != "" && lng != "" {
			location = pgtype.Text{String: fmt.Sprintf("POINT(%s %s)", lat, lng), Valid: true}
		} else {
			location = pgtype.Text{Valid: false}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		acc, err = p.accountService.GetAccountAction(ctx, account_id, user_id)
	}()
	wg.Wait()
	if err != nil {
		return res, err
	}

	descriptionNull := ConvertDescription(description)
	post, err := p.queries.CreatePost(ctx, db.CreatePostParams{
		PostTypeID:     1,
		AccountID:      account_id,
		Description:    descriptionNull,
		StGeomfromtext: location,
	})
	imgs := make([]db.PostImage, 0)
	for _, element := range images {
		images, _ := p.queries.AddImagePost(ctx, db.AddImagePostParams{PostID: post.ID, UrlImage: element})
		imgs = append(imgs, images)
	}

	res = PostRes(post, acc, imgs, 0)

	return res, nil
}

// DeleteImage implements IPostService.
func (p *PostService) DeleteImage(ctx context.Context, id int64) {
	panic("unimplemented")
}

// DeletePost implements IPostService.
func (p *PostService) DeletePost(ctx context.Context, id int64, user_id int64) error {
	panic("unimplemented")
}

// GetPost implements IPostService.
func (p *PostService) GetPost(ctx context.Context, id int64) (PostResponse, error) {
	var res PostResponse
	post, err := p.queries.GetPost(ctx, id)
	if err != nil {
		return res, err
	}
	acc, _ := p.queries.GetAccountById(ctx, post.AccountID)
	imgs, _ := p.queries.GetImagePost(ctx, post.ID)
	likes, _ := p.queries.CountReactPost(ctx, post.ID)
	res = GetPostRes(post, acc, imgs, likes)
	return res, nil
}

// UpdateContentPost implements IPostService.
func (p *PostService) UpdateContentPost(ctx context.Context, desciption string, id int64, user_id int64) (PostResponse, error) {
	panic("unimplemented")
}

func NewPostService(queries *db.Queries, account account.IAccountService) IPostService {
	return &PostService{
		queries:        queries,
		accountService: account,
	}
}
