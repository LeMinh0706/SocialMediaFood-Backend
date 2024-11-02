package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	*Queries
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTX(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("transaction exposure: %v, rollback err: %v", err, rbErr)
		}
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit failed: %v", err)
	}

	return nil
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=6,max=16"`
	Password string `json:"password" binding:"required,min=8"`
	Email    string `json:"email"`
	Fullname string `json:"fullname" binding:"required,min=6,max=32"`
	Gender   int32  `json:"gender" binding:"min=0,max=1"`
}

func (store *Store) CreateAccountTx(ctx context.Context, arg RegisterRequest) (RegisterRow, error) {
	var user RegisterRow
	hash, err := util.HashPassword(arg.Password)
	if err != nil {
		return user, err
	}
	var reqMail pgtype.Text
	if strings.TrimSpace(arg.Email) == "" {
		reqMail = pgtype.Text{Valid: false}
	} else {
		reqMail = pgtype.Text{String: arg.Email, Valid: true}
	}

	err = store.execTX(ctx, func(q *Queries) error {
		var err error
		user, err = q.Register(ctx, RegisterParams{
			Username:     arg.Username,
			Email:        reqMail,
			HashPassword: hash,
		})
		if err != nil {
			return err
		}

		_, err = q.CreateAccounts(ctx, CreateAccountsParams{
			UserID:               user.ID,
			Fullname:             arg.Fullname,
			Gender:               pgtype.Int4{Int32: arg.Gender, Valid: true},
			RoleID:               3,
			UrlAvatar:            util.RandomAvatar(arg.Gender),
			UrlBackgroundProfile: background,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return user, err
}
