package repo

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type CommentRepo struct {
	queries *db.Queries
}

func NewCommentRepo(queries *db.Queries) (*CommentRepo, error) {
	return &CommentRepo{
		queries: queries,
	}, nil
}

func (repo *CommentRepo) CreateComment(ctx context.Context, arg db.CreateCommentParams) (db.CreateCommentRow, error) {
	return repo.queries.CreateComment(ctx, arg)
}
