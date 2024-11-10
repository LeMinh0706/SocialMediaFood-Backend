package user

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type IUserService interface {
	Register(ctx context.Context, req db.RegisterRequest) (db.RegisterRow, error)
	Login(ctx context.Context, username, password string) (db.LoginRow, error)
	Backup(ctx context.Context)
}
