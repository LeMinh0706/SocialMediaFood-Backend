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

func (repo *ReactRepo) GetReact(ctx context.Context, arg db.GetReactParams) (int64, error) {
	return repo.queries.GetReact(ctx, arg)
}

func (repo *ReactRepo) GetReactPost(ctx context.Context, page, pageSize int32, post_id int64) ([]db.GetReactPostRow, error) {
	return repo.queries.GetReactPost(ctx, db.GetReactPostParams{
		PostID: post_id,
		Limit:  pageSize,
		Offset: (page - 1) * pageSize,
	})
}

func (repo *ReactRepo) CountLike(ctx context.Context, post_id int64) (int64, error) {
	return repo.queries.CountReactPost(ctx, post_id)
}

func (repo *ReactRepo) UpdateState(ctx context.Context, post_id, account_id int64, state int32) (db.ReactPost, error) {
	return repo.queries.UpdateState(ctx, db.UpdateStateParams{PostID: post_id, AccountID: account_id, State: state})
}

func (repo *ReactRepo) UnlikePost(ctx context.Context, post_id, account_id int64) error {
	return repo.queries.DeleteReact(ctx, db.DeleteReactParams{PostID: post_id, AccountID: account_id})
}
