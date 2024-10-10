package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type PostRepository struct {
	queries *db.Queries
}

func NewPostRepo(queries *db.Queries) (*PostRepository, error) {
	return &PostRepository{
		queries: queries,
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

func (repo *PostRepository) CreateImagePost(ctx context.Context, post_id int64, imageUrl string) (db.PostImage, error) {
	return repo.queries.CreateImagePost(ctx, db.CreateImagePostParams{
		PostID:   post_id,
		UrlImage: imageUrl,
	})
}

func (repo *PostRepository) GetImagePost(ctx context.Context, post_id int64) ([]db.PostImage, error) {
	return repo.queries.GetImagePost(ctx, post_id)
}

func (repo *PostRepository) GetPost(ctx context.Context, id int64) (db.Post, error) {
	return repo.queries.GetPost(ctx, id)
}

func (repo *PostRepository) GetListPost(ctx context.Context, page, pageSize int32) ([]db.Post, error) {
	return repo.queries.ListPost(ctx, db.ListPostParams{
		Limit:  pageSize,
		Offset: page*pageSize - (pageSize - 1),
	})
}
