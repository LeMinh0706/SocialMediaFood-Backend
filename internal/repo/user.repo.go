package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/db"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepo() (*UserRepository, error) {
	pg, err := getDBConnection()
	if err != nil {
		return nil, err
	}
	return &UserRepository{
		queries: db.New(pg),
	}, nil
}

func (repo *UserRepository) CreateUser(ctx context.Context, username, password string, role_id int32) (db.User, error) {
	return repo.queries.CreateUser(ctx, db.CreateUserParams{
		Username:          username,
		Fullname:          username,
		HashPashword:      password,
		Email:             sql.NullString{Valid: false},
		Gender:            0,
		RoleID:            role_id,
		DateCreateAccount: time.Now().Unix(),
	})
}

func (repo *UserRepository) GetUser(ctx context.Context, username string) (db.User, error) {
	return repo.queries.GetUser(ctx, username)
}

func (repo *UserRepository) GetUserById(ctx context.Context, id int64) (db.GetUserByIdRow, error) {
	return repo.queries.GetUserById(ctx, id)
}
