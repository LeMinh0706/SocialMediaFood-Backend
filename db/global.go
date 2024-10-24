package db

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/jackc/pgx/v5"
)

var (
	pgd        *pgx.Conn
	background = "upload/background/background_1.jpg"
)

// Get connection để tương tác dưới db, ở repo khỏi phải viết lại nguyên hàm
func GetDBConnection() (*pgx.Conn, error) {
	if pgd == nil {
		config, err := util.LoadConfig("../..")
		if err != nil {
			return nil, err
		}
		pgd, err = pgx.Connect(context.Background(), config.DBSource)
		if err != nil {
			return nil, err
		}
	}
	return pgd, nil
}

func GetBackground() string {
	return background
}
