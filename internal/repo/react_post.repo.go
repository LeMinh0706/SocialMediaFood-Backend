package repo

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type ReactPostRepository struct {
	queries *db.Queries
}

func NewReactPostRepo(queries *db.Queries) (*ReactPostRepository, error) {
	return &ReactPostRepository{
		queries: queries,
	}, nil
}

func (repo *ReactPostRepository) CreateReact(ctx context.Context, arg db.CreateReactParams) (db.ReactPost, error) {
	return repo.queries.CreateReact(ctx, arg)
}

func (repo *ReactPostRepository) GetReact(ctx context.Context, arg db.GetReactParams) (db.ReactPost, error) {
	return repo.queries.GetReact(ctx, arg)
}

// /Get all react in post
func (repo *ReactPostRepository) GetPostReact(ctx context.Context, post_id int64) ([]int64, error) {
	return repo.queries.GetReactPost(ctx, post_id)
}

// func (repo *ReactPostRepository) CountReact(ctx context.Context, post_id int64) (int64, error) {
// 	return repo.queries.GetCountPost()
// }

///

func (repo *ReactPostRepository) DeleteReact(ctx context.Context, id int64) error {
	return repo.queries.DeleteReact(ctx, id)
}
