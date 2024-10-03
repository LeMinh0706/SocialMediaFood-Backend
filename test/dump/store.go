package dump

import (
	"context"
	"database/sql"
	"fmt"

	d "github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type Store struct {
	*d.Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: d.New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*d.Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := d.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("Transaction error: %v - Rollback error: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type PostTxParams struct {
	Id             int64          `json:"id"`
	PostTypeID     int32          `json:"post_type_id"`
	UserID         int64          `json:"user_id"`
	Description    sql.NullString `json:"description"`
	DateCreatePost int64          `json:"date_create_post"`
}

type PostTxResult struct {
	Post        d.Post         `json:"post"`
	Id          int64          `json:"id"`
	PostTypeID  int32          `json:"post_type_id"`
	UserID      int64          `json:"user_id"`
	Description sql.NullString `json:"description"`
	Image       d.PostImage    `json:"image_post"`
	PostId      int64          `json:"post_id"`
}

func (store *Store) CreatePostTx(ctx context.Context, arg PostTxParams, imageName []string) (PostTxResult, error) {
	var result PostTxResult
	err := store.execTx(ctx, func(q *d.Queries) error {
		var err error
		result.Post, err = q.CreatePost(ctx, d.CreatePostParams{
			PostTypeID:     arg.PostTypeID,
			UserID:         arg.UserID,
			Description:    arg.Description,
			DateCreatePost: arg.DateCreatePost,
		})
		if err != nil {
			return err
		}
		for _, imgUrl := range imageName {
			result.Image, err = q.CreateImagePost(ctx, d.CreateImagePostParams{
				UrlImage: imgUrl,
				PostID:   result.Id,
			})
		}
		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}
