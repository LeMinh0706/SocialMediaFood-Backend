package reset_password

import (
	"context"
	"fmt"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/mails"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type ResetPasswordService struct {
	queries *db.Queries
}

// SendMail implements IResetPasswordService.
func (r *ResetPasswordService) SendMail(ctx context.Context, email string, token string, config util.Config) error {
	link := fmt.Sprintf("%v?token=%v", config.FrontEndUrl, token)
	err := mails.SendMailResetPassword([]string{email}, link, config)
	if err != nil {
		return err
	}
	return nil
}

// ForgotPassword implements IResetPasswordService.
func (r *ResetPasswordService) ForgotPassword(ctx context.Context, email string, config util.Config) (db.GetUserByEmailRow, error) {
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

// AddRequestPassword implements IResetPasswordService.
func (r *ResetPasswordService) AddRequestPassword(ctx context.Context, id uuid.UUID, user_id int64, duration time.Duration) error {
	_, err := r.queries.CreateRequestPassword(ctx, db.CreateRequestPasswordParams{
		ID:        pgtype.UUID{Bytes: id, Valid: true},
		UserID:    user_id,
		ExpiresAt: pgtype.Timestamptz{Time: time.Now().Add(duration), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

// ChangePassword implements IResetPasswordService.
func (r *ResetPasswordService) ChangePassword(ctx context.Context, uuid uuid.UUID, user_id int64, password string) error {
	check, err := r.queries.GetRequestByUUID(ctx, pgtype.UUID{Bytes: uuid, Valid: true})
	if err != nil {
		return fmt.Errorf("invalid uuid")
	}
	if check.IsActive {
		return fmt.Errorf("request used")
	}
	hash, _ := util.HashPassword(password)
	err = r.queries.UpdateActive(ctx, pgtype.UUID{Bytes: uuid, Valid: true})
	if err != nil {
		return err
	}
	_, err = r.queries.UpdatePassword(ctx, db.UpdatePasswordParams{ID: user_id, HashPassword: hash})
	if err != nil {
		return err
	}
	return nil
}

func NewResetPasswordService(queries *db.Queries) IResetPasswordService {
	return &ResetPasswordService{
		queries: queries,
	}
}
