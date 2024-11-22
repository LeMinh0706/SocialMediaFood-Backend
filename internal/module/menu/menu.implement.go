package menu

import (
	"context"
	"math/big"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type MenuService struct {
	queries *db.Queries
}

// CreateDish implements IMenuService.
func (m *MenuService) CreateDish(ctx context.Context, account_id int64, quantity int32, name string, img string, price float64) (db.Menu, error) {
	// var res MenuResponse
	dish, err := m.queries.AddToMenu(ctx, db.AddToMenuParams{
		AccountID: account_id,
		DishName:  name,
		Quantity:  quantity,
		Price: pgtype.Numeric{
			Exp:   -3,
			Int:   big.NewInt(int64(price * 1000)),
			Valid: true,
		},
		Img: img,
	})
	if err != nil {
		return db.Menu{}, err
	}
	return dish, err
}

// GetMenu implements IMenuService.
func (m *MenuService) GetMenu(ctx context.Context, account_id int64, page int32, pageSize int32) ([]db.Menu, error) {
	list, err := m.queries.GetMenu(ctx, db.GetMenuParams{
		AccountID: account_id,
		Limit:     pageSize,
		Offset:    (page - 1) * pageSize,
	})
	if err != nil {
		return nil, err
	}
	return list, err
}

// DeleteFood implements IMenuService.
func (m *MenuService) DeleteFood(ctx context.Context, id int64) error {
	panic("unimplemented")
}

func NewMenuService(queries *db.Queries) IMenuService {
	return &MenuService{
		queries: queries,
	}
}
