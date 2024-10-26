package db

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	background = "upload/background/background_1.jpg"
)

// Get connection để tương tác dưới db, ở repo khỏi phải viết lại nguyên hàm
// pgx nhanh hơn, nhưng xử phải có Maxconn hợp lí để ko nghẽn cổ chai
func GetDBConnection(config util.Config) (*pgxpool.Pool, error) {

	pgConfig, err := pgxpool.ParseConfig(config.DBSource)
	if err != nil {
		return nil, err
	}

	pgConfig.MaxConns = config.MaxConns
	pgConfig.MinConns = config.MinConns
	pgConfig.MaxConnLifetime = config.MaxConnLifetime
	pgConfig.MaxConnIdleTime = config.MaxConnIdleTime

	pgd, err := pgxpool.NewWithConfig(context.Background(), pgConfig)
	if err != nil {
		return nil, err
	}
	return pgd, nil
}

func GetBackground() string {
	return background
}
