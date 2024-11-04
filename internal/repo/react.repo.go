package repo

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type ReactRepo struct {
	queries *db.Queries
}

func NewReactRepo(queries *db.Queries) (*ReactRepo, error) {
	return &ReactRepo{
		queries: queries,
	}, nil
}

func (repo *ReactRepo) CreateReact(ctx context.Context, arg db.CreateReactParams) (db.ReactPost, error) {
	return repo.queries.CreateReact(ctx, arg)
}

func (repo *ReactRepo) GetReactPost(ctx context.Context, page, pageSize int32, post_id int64) ([]db.GetReactPostRow, error) {
	return repo.queries.GetReactPost(ctx, db.GetReactPostParams{
		PostID: post_id,
		Limit:  pageSize,
		Offset: (page - 1) * pageSize,
	})
}
