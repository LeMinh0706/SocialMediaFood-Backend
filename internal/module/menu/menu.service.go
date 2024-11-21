package menu

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type IMenuService interface {
	CreateDish(ctx context.Context, account_id int64, quantity int32, name, img string, price float64) (db.Menu, error)
	GetMenu(ctx context.Context, account_id int64, page, pageSize int32) ([]db.Menu, error)
	DeleteFood(ctx context.Context, id int64) error
}
