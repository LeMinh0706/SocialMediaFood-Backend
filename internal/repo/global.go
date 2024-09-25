package repo

import (
	"database/sql"

	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
)

var pgd *sql.DB

func getDBConnection() (*sql.DB, error) {
	if pgd == nil {
		config, err := util.LoadConfig("../..") // Hoặc bất kỳ nơi nào bạn load config
		if err != nil {
			return nil, err
		}
		pgd, err = sql.Open(config.DBDriver, config.DBSource)
		if err != nil {
			return nil, err
		}
		if err := pgd.Ping(); err != nil {
			return nil, err
		}
	}
	return pgd, nil
}
