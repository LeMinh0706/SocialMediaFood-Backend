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

func (repo *UserRepository) CreateUser(fullname, password string) (db.User, error) {
	return repo.queries.CreateUser(context.Background(), db.CreateUserParams{
		Fullname:          fullname,
		HashPashword:      password,
		Email:             sql.NullString{Valid: false},
		Gender:            0,
		RoleID:            3,
		DateCreateAccount: time.Now().Unix(),
	})
}
