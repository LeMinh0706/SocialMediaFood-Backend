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
