package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepo() (*UserRepository, error) {
	pg, err := db.GetDBConnection()
	if err != nil {
		return nil, err
	}
	return &UserRepository{
		queries: db.New(pg),
	}, nil
}

func (repo *UserRepository) CreateUser(ctx context.Context, username, password string, gender int32, role_id int32) (db.User, error) {
	return repo.queries.CreateUser(ctx, db.CreateUserParams{
		Username:             username,
		Fullname:             username,
		HashPashword:         password,
		Email:                sql.NullString{Valid: false},
		Gender:               gender,
		UrlAvatar:            util.RandomAvatar(gender),
		UrlBackgroundProfile: db.GetBackground(),
		RoleID:               role_id,
		DateCreateAccount:    time.Now().Unix(),
	})
}

func (repo *UserRepository) GetUser(ctx context.Context, username string) (db.User, error) {
	return repo.queries.GetUser(ctx, username)
}

func (repo *UserRepository) GetUserById(ctx context.Context, id int64) (db.GetUserByIdRow, error) {
	return repo.queries.GetUserById(ctx, id)
}
