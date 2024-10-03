package post

import (
	"context"
	"database/sql"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type PostRepository struct {
	queries *db.Queries
}

func NewPostRepo() (*PostRepository, error) {
	pg, err := db.GetDBConnection()
	if err != nil {
		return nil, err
	}

	return &PostRepository{
		queries: db.New(pg),
	}, nil
}

func (repo *PostRepository) CreatePost(ctx context.Context, description string, user_id int64) (db.Post, error) {
	return repo.queries.CreatePost(ctx, db.CreatePostParams{
		PostTypeID:     1,
		UserID:         user_id,
		Description:    sql.NullString{String: description, Valid: true},
		DateCreatePost: time.Now().Unix(),
	})
}
