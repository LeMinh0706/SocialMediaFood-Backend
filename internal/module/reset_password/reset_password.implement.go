package reset_password

import (
	"context"
	"fmt"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type ResetPasswordService struct {
	queries *db.Queries
}

// ChangePassword implements IResetPasswordService.
func (r *ResetPasswordService) ChangePassword(ctx context.Context, uuid uuid.UUID, user_id int64, username string, password string) error {
	panic("unimplemented")
}

// ResetPasswordRequest implements IResetPasswordService.
func (r *ResetPasswordService) ResetPasswordRequest(ctx context.Context, email string) (db.GetUserByEmailRow, error) {
	var res db.GetUserByEmailRow
	user, err := r.queries.GetUserByEmail(ctx, pgtype.Text{String: email, Valid: true})
	if err != nil {
		return res, fmt.Errorf("not found")
	}
	check, err := r.queries.GetCheckAction(ctx, user.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return user, nil
		}
		return res, err
	}

	if check.IsActive {
		return user, nil
	}

	if !time.Now().After(check.ExpiresAt.Time) {
		return res, fmt.Errorf("intime")
	}
	return user, nil
}

// SendMail implements IResetPasswordService.
func (r *ResetPasswordService) SendMail(ctx context.Context, arg db.GetUserByEmailRow) error {
	panic("unimplemented")
}

func NewResetPasswordService(queries *db.Queries) IResetPasswordService {
	return &ResetPasswordService{
		queries: queries,
	}
}
