package menu

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type IMenuService interface {
	CreateDish(ctx context.Context, arg db.AddToMenuParams) (db.Menu, error)
	GetMenu(ctx context.Context, arg db.GetMenuParams) ([]db.GetMenuRow, error)
	DeleteFood(ctx context.Context, id int64) error
}
