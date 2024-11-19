package menu

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type MenuService struct {
	queries *db.Queries
}

// CreateDish implements IMenuService.
func (m *MenuService) CreateDish(ctx context.Context, arg db.AddToMenuParams) (db.Menu, error) {
	panic("unimplemented")
}

// DeleteFood implements IMenuService.
func (m *MenuService) DeleteFood(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// GetMenu implements IMenuService.
func (m *MenuService) GetMenu(ctx context.Context, arg db.GetMenuParams) ([]db.GetMenuRow, error) {
	panic("unimplemented")
}

func NewMenuService(queries *db.Queries) IMenuService {
	return &MenuService{
		queries: queries,
	}
}
