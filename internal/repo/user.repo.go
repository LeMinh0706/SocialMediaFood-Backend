package repo

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepo struct {
	queries *db.Queries
}

func NewUserRepo(queries *db.Queries) (*UserRepo, error) {
	return &UserRepo{
		queries: queries,
	}, nil
}

func (repo *UserRepo) Register(ctx context.Context, username, password string, email pgtype.Text) (db.RegisterRow, error) {
	return repo.queries.Register(ctx, db.RegisterParams{
		Username:     username,
		Email:        email,
		HashPassword: password,
	})
}

func (repo *UserRepo) Login(ctx context.Context, username string) (db.LoginRow, error) {
	return repo.queries.Login(ctx, username)
}
