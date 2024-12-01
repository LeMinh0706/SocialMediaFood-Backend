package reset_password

import (
	"context"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/google/uuid"
)

type IResetPasswordService interface {
	ForgotPassword(ctx context.Context, email string, config util.Config) (db.GetUserByEmailRow, error)
	AddRequestPassword(ctx context.Context, id uuid.UUID, user_id int64, duration time.Duration) error
	SendMail(ctx context.Context, email string, token string, config util.Config) error
	ChangePassword(ctx context.Context, uuid uuid.UUID, user_id int64, password string) error
	SpamMail(ctx context.Context, email, name string, config util.Config) error
}
