package repo

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/jackc/pgx/v5/pgtype"
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

func (repo *CommentRepo) GetComment(ctx context.Context, id int64) (db.GetCommentRow, error) {
	return repo.queries.GetComment(ctx, id)
}

func (repo *CommentRepo) GetListComment(ctx context.Context, page, pageSize int32, post_id int64) ([]int64, error) {
	return repo.queries.GetListComment(ctx, db.GetListCommentParams{
		PostTopID: pgtype.Int8{Int64: post_id, Valid: true},
		Limit:     pageSize,
		Offset:    (page - 1) * pageSize,
	})
}
