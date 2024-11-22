package menu

import "github.com/LeMinh0706/SocialMediaFood-Backend/db"

type AddFood struct {
	AccountID int64   `json:"account_id"`
	DishName  string  `json:"dish_name"`
	Quantity  int32   `json:"quantity"`
	Price     float64 `json:"price"`
	Img       string  `json:"img"`
}

type MenuResponse struct {
	ID        int64   `json:"id"`
	AccountID int64   `json:"account_id"`
	DishName  string  `json:"dish_name"`
	Quantity  int32   `json:"quantity"`
	Price     float64 `json:"price"`
	Img       string  `json:"img"`
	IsDelete  bool    `json:"is_delete"`
}

func MenuRes(dish db.Menu) MenuResponse {
	priceFloat, _ := dish.Price.Float64Value()
	return MenuResponse{
		ID:        dish.ID,
		AccountID: dish.AccountID,
		DishName:  dish.DishName,
		Quantity:  dish.Quantity,
		Img:       dish.Img,
		IsDelete:  dish.IsDelete,
		Price:     priceFloat.Float64,
	}
}
