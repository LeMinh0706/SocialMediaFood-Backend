package rating

import (
	"context"
	"fmt"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/account"
	"github.com/jackc/pgx/v5/pgtype"
)

type RatingService struct {
	queries *db.Queries
	acc     account.IAccountService
}

// CreateRating implements IRatingService.
func (r *RatingService) CreateRating(ctx context.Context, req RatingRequest) error {
	acc, err := r.queries.GetAccountById(ctx, req.ToAccountID)
	if err != nil {

		return fmt.Errorf("it's not owner")
	}
	if acc.RoleID != 2 {
		return fmt.Errorf("it's not owner")
	}
	arg := db.CreateRatingParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Star:          req.Star,
		Content:       pgtype.Text{String: req.Content, Valid: true},
	}
	_, err = r.queries.CreateRating(ctx, arg)
	if err != nil {
		return fmt.Errorf("can't rate this owner")
	}
	return nil
}

// DeleteRating implements IRatingService.
func (r *RatingService) DeleteRating(ctx context.Context, from_account_id int64, to_account_id int64) error {
	panic("unimplemented")
}

// GetListRating implements IRatingService.
func (r *RatingService) GetListRating(ctx context.Context, account_id int64, page int32, page_size int32) ([]db.Rating, error) {
	panic("unimplemented")
}

// UpdateRating implements IRatingService.
func (r *RatingService) UpdateRating(ctx context.Context, from_account_id int64, to_account_id int64, content string, start int32) error {
	panic("unimplemented")
}

func NewRatingService(q *db.Queries, a account.IAccountService) IRatingService {
	return &RatingService{
		queries: q,
		acc:     a,
	}
}
