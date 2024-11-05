package models

import (
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type PostResponse struct {
	ID          int64          `json:"id"`
	PostTypeID  int32          `json:"post_type_id"`
	AccountID   int64          `json:"account_id"`
	Description string         `json:"description"`
	Lng         interface{}    `json:"lng"`
	Lat         interface{}    `json:"lat"`
	CreatedAt   time.Time      `json:"created_at"`
	Images      []db.PostImage `json:"images"`
	Account     AccountForPost `json:"account"`
	TotalLike   int64          `json:"total_like"`
}

type UpdatePostRequest struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
}

func PostRes(post db.CreatePostRow, account AccountForPost, imgs []db.PostImage, totalLike int64) PostResponse {
	return PostResponse{
		ID:          post.ID,
		PostTypeID:  post.PostTypeID,
		AccountID:   post.AccountID,
		Description: post.Description.String,
		Lng:         post.Lng,
		Lat:         post.Lat,
		CreatedAt:   post.CreatedAt.Time,
		Images:      imgs,
		Account:     account,
		TotalLike:   totalLike,
	}
}

func UpdatePostRes(post db.UpdatePostRow, account AccountForPost, imgs []db.PostImage) PostResponse {
	return PostResponse{
		ID:          post.ID,
		PostTypeID:  post.PostTypeID,
		AccountID:   post.AccountID,
		Description: post.Description.String,
		Lng:         post.Lng,
		Lat:         post.Lat,
		CreatedAt:   post.CreatedAt.Time,
		Images:      imgs,
		Account:     account,
	}
}
