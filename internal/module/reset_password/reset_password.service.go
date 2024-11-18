package reset_password

import (
	"context"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/google/uuid"
)

type IResetPasswordService interface {
	ForgotPassword(ctx context.Context, email string) (db.GetUserByEmailRow, error)
	AddRequestPassword(ctx context.Context, id uuid.UUID, user_id int64, duration time.Duration) error
	SendMail(ctx context.Context, arg db.GetUserByEmailRow) error
	ChangePassword(ctx context.Context, uuid uuid.UUID, user_id int64, password string) error
}
