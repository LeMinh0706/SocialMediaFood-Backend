package admin

import "time"

type UpgradePrice struct {
	ID        int32     `json:"id"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
