package rating

import (
	"time"
)

type RatingRequest struct {
	FromAccountID int64  `json:"from_account_id"`
	ToAccountID   int64  `json:"to_account_id"`
	Star          int32  `json:"star" binding:"required,min=1,max=5"`
	Content       string `json:"content"`
}

type DeleteRatingRequest struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
}

type ListRating struct {
	FromAccountID        int64     `json:"from_account_id"`
	ToAccountID          int64     `json:"to_account_id"`
	Star                 int32     `json:"star"`
	Content              string    `json:"content"`
	CreatedAt            time.Time `json:"created_at"`
	Fullname             string    `json:"fullname"`
	UrlAvatar            string    `json:"url_avatar"`
	UrlBackgroundProfile string    `json:"url_background_profile"`
}
