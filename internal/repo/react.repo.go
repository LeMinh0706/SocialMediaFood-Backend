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
