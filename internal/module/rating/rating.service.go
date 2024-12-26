package rating

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type IRatingService interface {
	CreateRating(ctx context.Context, req RatingRequest) error
	GetListRating(ctx context.Context, account_id int64, page, page_size int32) ([]db.Rating, error)
	UpdateRating(ctx context.Context, from_account_id, to_account_id int64, content string, start int32) error
	DeleteRating(ctx context.Context, from_account_id, to_account_id int64) error
}
