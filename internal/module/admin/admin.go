package admin

import (
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/post"
)

type UpgradePrice struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Benefit   string    `json:"benefit" binding:"required"`
	Price     float64   `json:"price"`
	IsChoose  bool      `json:"is_choose"`
	CreatedAt time.Time `json:"created_at"`
}

type AddUpgradePrice struct {
	Title   string  `json:"title" binding:"required"`
	Benefit string  `json:"benefit" binding:"required"`
	Price   float64 `json:"price" binding:"required"`
}

type ReportFrom struct {
	Account db.GetAccountByIdRow `json:"account"`
	Issue   db.IssuePost         `json:"issue"`
}

type ReportDetailResponse struct {
	Post       post.PostResponse `json:"posts"`
	IssuePosts []ReportFrom      `json:"report_from"`
}
