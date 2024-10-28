package repo

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type UserRepo struct {
	queries *db.Queries
	store   *db.Store
}

func NewUserRepo(queries *db.Queries, store *db.Store) (*UserRepo, error) {
	return &UserRepo{
		queries: queries,
		store:   store,
	}, nil
}

func (repo *UserRepo) RegisterTx(ctx context.Context, arg db.RegisterRequest) (db.RegisterRow, error) {
	return repo.store.CreateAccountTx(ctx, arg)
}

func (repo *UserRepo) Login(ctx context.Context, username string) (db.LoginRow, error) {
	return repo.queries.Login(ctx, username)
}
