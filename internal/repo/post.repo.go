package repo

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type PostRepo struct {
	queries *db.Queries
}

func NewPostRepo(queries *db.Queries) (*PostRepo, error) {
	return &PostRepo{
		queries: queries,
	}, nil
}

func (repo *PostRepo) CreatePost(ctx context.Context, arg db.CreatePostParams) (db.CreatePostRow, error) {
	return repo.queries.CreatePost(ctx, db.CreatePostParams{
		PostTypeID:     arg.PostTypeID,
		AccountID:      arg.AccountID,
		Description:    arg.Description,
		StGeomfromtext: arg.StGeomfromtext,
	})
}

func (repo *PostRepo) CreateImagePost(ctx context.Context, arg db.AddImagePostParams) (db.PostImage, error) {
	return repo.queries.AddImagePost(ctx, arg)
}

func (repo *PostRepo) GetImagePost(ctx context.Context, post_id int64) ([]db.PostImage, error) {
	return repo.queries.GetImagePost(ctx, post_id)
}

func (repo *PostRepo) GetListPost(ctx context.Context, page, pageSize int32) ([]int64, error) {
	return repo.queries.GetListPost(ctx, db.GetListPostParams{Limit: pageSize, Offset: (page - 1) * pageSize})
}

func (repo *PostRepo) GetPost(ctx context.Context, id int64) (db.GetPostRow, error) {
	return repo.queries.GetPost(ctx, id)
}

func (repo *PostRepo) DeletePost(ctx context.Context, id int64) error {
	return repo.queries.DeletePost(ctx, id)
}
