package post

import (
	"context"
	"fmt"
	"sync"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/account"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type PostService struct {
	queries        *db.Queries
	accountService account.IAccountService
}

// DeleteImage implements IPostService.
func (p *PostService) DeleteImage(ctx context.Context, username string, id int64) error {
	image, _ := p.queries.GetImage(ctx, id)
	post, _ := p.GetPost(ctx, 0, image.PostID)
	_, err := p.accountService.GetAccountAction(ctx, post.AccountID, username)
	if err != nil {
		return err
	}
	err = p.queries.DeleteImagePost(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// DeletePost implements IPostService.
func (p *PostService) DeletePost(ctx context.Context, username string, id int64) error {
	post, err := p.GetPost(ctx, 0, id)
	if err != nil {
		return err
	}
	_, err = p.accountService.GetAccountAction(ctx, post.AccountID, username)
	if err != nil {
		return err
	}
	err = p.queries.DeletePost(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateContentPost implements IPostService.
func (p *PostService) UpdateContentPost(ctx context.Context, username string, desciption string, id int64, images []string) (PostResponse, error) {
	var res PostResponse
	post, err := p.GetPost(ctx, 0, id)
	if err != nil {
		return res, err
	}
	acc, err := p.accountService.GetAccountAction(ctx, post.AccountID, username)
	if err != nil {
		return res, err
	}

	for _, element := range images {
		p.queries.AddImagePost(ctx, db.AddImagePostParams{PostID: post.ID, UrlImage: element})
	}

	update, err := p.queries.UpdatePost(ctx, db.UpdatePostParams{
		ID:          id,
		Description: ConvertDescription(desciption),
	})

	if err != nil {
		return res, err
	}

	pic, _ := p.queries.GetImagePost(ctx, id)

	res = PostRes(db.CreatePostRow(update), acc, pic, post.ReactState, post.TotalLike, post.TotalComment)
	return res, nil
}

// CreatePost implements IPostService.
func (p *PostService) CreatePost(ctx context.Context, username string, description string, lng string, lat string, images []string, account_id int64) (PostResponse, error) {
	var res PostResponse
	var wg sync.WaitGroup
	var location pgtype.Text
	var acc db.Account
	var err error
	wg.Add(1)
	go func() {
		defer wg.Done()
		if lat != "" && lng != "" {
			location = pgtype.Text{String: fmt.Sprintf("POINT(%s %s)", lng, lat), Valid: true}
		} else {
			location = pgtype.Text{Valid: false}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		acc, err = p.accountService.GetAccountAction(ctx, account_id, username)
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

	res = PostRes(post, acc, imgs, db.ReactPost{AccountID: account_id, PostID: post.ID, State: 0}, 0, 0)

	return res, nil
}

// GetListImage implements IPostService.
func (p *PostService) GetListImage(ctx context.Context, page int32, pageSize int32) ([]db.PostImage, error) {
	list, err := p.queries.GetListImage(ctx, db.GetListImageParams{
		Limit:  pageSize,
		Offset: (page - 1) * pageSize,
	})
	if err != nil {
		return []db.PostImage{}, err
	}
	return list, err
}

// GetPostInLocate implements IPostService.
func (p *PostService) GetPostInLocate(ctx context.Context, dwithin int64, account_id int64, lng string, lat string, page int32, pageSize int32) ([]PostResponse, error) {
	var res []PostResponse
	list, err := p.queries.GetPostInLocate(ctx, db.GetPostInLocateParams{
		StGeomfromtext: fmt.Sprintf("POINT(%s %s)", lng, lat),
		StDwithin:      dwithin,
		Limit:          pageSize,
		Offset:         (page - 1) * pageSize,
	})
	if err != nil {
		return res, nil
	}
	for _, element := range list {
		post, _ := p.GetPost(ctx, account_id, element)
		res = append(res, post)
	}
	return res, nil
}

// GetImage implements IPostService.
func (p *PostService) GetImage(ctx context.Context, id int64) ([]db.PostImage, error) {
	images, err := p.queries.GetImagePost(ctx, id)
	if err != nil {
		return nil, err
	}
	return images, nil
}

// GetHomePagePost implements IPostService.
func (p *PostService) GetHomePagePost(ctx context.Context, account_id int64, page int32, pageSize int32) ([]PostResponse, error) {
	var res []PostResponse
	ps := pageSize * 4
	list, err := p.queries.GetHomePagePost(ctx, db.GetHomePagePostParams{
		AccountID: account_id,
		Limit:     ps,
		Offset:    (page - 1) * ps,
	})
	if err != nil {
		return res, err
	}
	for _, element := range list {
		post, err := p.GetPost(ctx, account_id, element)
		if err != nil {
			return res, err
		}
		res = append(res, post)
	}
	return res, nil
}

// GetListPost implements IPostService.
func (p *PostService) GetListPost(ctx context.Context, page int32, pageSize int32) ([]PostResponse, error) {
	var res []PostResponse
	list, err := p.queries.GetListPost(ctx, db.GetListPostParams{Limit: pageSize, Offset: (page - 1) * pageSize})
	if err != nil {
		return res, err
	}
	for _, element := range list {
		post, err := p.GetPost(ctx, 0, element)
		if err != nil {
			return []PostResponse{}, err
		}
		res = append(res, post)
	}
	return res, nil
}

// GetPersonPost implements IPostService.
func (p *PostService) GetPersonPost(ctx context.Context, from int64, to int64, page int32, pageSize int32) ([]PostResponse, error) {
	var res []PostResponse
	list, err := p.queries.GetPersonPost(ctx, db.GetPersonPostParams{AccountID: to, Limit: pageSize, Offset: (page - 1) * pageSize})
	if err != nil {
		return res, err
	}
	for _, element := range list {
		post, err := p.GetPost(ctx, from, element)
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

// GetPost implements IPostService.
func (p *PostService) GetPost(ctx context.Context, account_id int64, id int64) (PostResponse, error) {
	var res PostResponse
	var wg sync.WaitGroup
	var acc db.GetAccountByIdRow
	var imgs []db.PostImage
	var countComment int64
	var react db.ReactPost
	var countReact int64
	errChan := make(chan error, 5)
	post, err := p.queries.GetPost(ctx, id)
	if err != nil {
		return res, err
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		acc, _ = p.queries.GetAccountById(ctx, post.AccountID)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		countReact, _ = p.queries.CountReactPost(ctx, post.ID)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		countComment, _ = p.queries.CountComment(ctx, pgtype.Int8{Int64: id, Valid: true})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		imgs, _ = p.queries.GetImagePost(ctx, post.ID)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		react, err = p.queries.GetReact(ctx, db.GetReactParams{
			AccountID: account_id,
			PostID:    id,
		})
		if err != nil {
			if err == pgx.ErrNoRows {
				react = db.ReactPost{AccountID: account_id, PostID: id, State: 0}
			} else {
				errChan <- err
			}
		}

	}()
	go func() {
		wg.Wait()
		close(errChan)
	}()
	for err := range errChan {
		if err != nil {
			return res, err
		}
	}
	res = GetPostRes(post, acc, imgs, react, countReact, countComment)
	return res, nil
}

func NewPostService(queries *db.Queries, account account.IAccountService) IPostService {
	return &PostService{
		queries:        queries,
		accountService: account,
	}
}
