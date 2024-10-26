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
		PostTypeID:    arg.PostTypeID,
		AccountID:     arg.AccountID,
		Description:   arg.Description,
		StMakepoint:   arg.StMakepoint,
		StMakepoint_2: arg.StMakepoint_2,
	})
}
