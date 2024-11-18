package reset_password

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/google/uuid"
)

type IResetPasswordService interface {
	ResetPasswordRequest(ctx context.Context, email string) (db.GetUserByEmailRow, error)
	SendMail(ctx context.Context, arg db.GetUserByEmailRow) error
	ChangePassword(ctx context.Context, uuid uuid.UUID, user_id int64, username string, password string) error
}
