package comment

import (
	"context"
	"database/sql"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type CommentRepository struct {
	queries *db.Queries
}

func NewCommentRepo() (*CommentRepository, error) {
	pg, err := db.GetDBConnection()
	if err != nil {
		return nil, err
	}
	return &CommentRepository{
		queries: db.New(pg),
	}, nil
}

func (c *CommentRepository) CreateComment(ctx context.Context, description string, user_id, post_top_id int64) (db.Post, error) {
	return c.queries.CreateComment(ctx, db.CreateCommentParams{
		UserID:         user_id,
		PostTopID:      sql.NullInt64{Int64: post_top_id, Valid: true},
		Description:    sql.NullString{String: description, Valid: true},
		DateCreatePost: time.Now().Unix(),
	})
}

func (c *CommentRepository) ListComment(ctx context.Context, post_id int64, page, pageSize int32) ([]db.Post, error) {
	return c.queries.ListComment(ctx, db.ListCommentParams{
		PostTopID: sql.NullInt64{Int64: post_id, Valid: true},
		Limit:     pageSize,
		Offset:    page*pageSize - (pageSize - 1),
	})
}

func (c *CommentRepository) GetCommentById(ctx context.Context, id int64) (db.Post, error) {
	return c.queries.GetCommentById(ctx, id)
}

func (c *CommentRepository) UpdateComment(ctx context.Context, id int64, description string) (db.Post, error) {
	return c.queries.UpdateComment(ctx, db.UpdateCommentParams{
		ID:          id,
		Description: sql.NullString{String: description, Valid: true},
	})
}
