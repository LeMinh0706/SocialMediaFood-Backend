package post

import (
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type PostResponse struct {
	ID           int64                `json:"id"`
	PostTypeID   int32                `json:"post_type_id"`
	AccountID    int64                `json:"account_id"`
	Description  string               `json:"description"`
	Lng          interface{}          `json:"lng"`
	Lat          interface{}          `json:"lat"`
	CreatedAt    time.Time            `json:"created_at"`
	Images       []db.PostImage       `json:"images"`
	Account      db.GetAccountByIdRow `json:"account"`
	ReactState   db.ReactPost         `json:"react_state"`
	TotalLike    int64                `json:"total_like"`
	TotalComment int64                `json:"total_comment"`
}

func PostRes(post db.CreatePostRow, account db.GetAccountByIdRow, imgs []db.PostImage, reactState db.ReactPost, totalLike, totalComment int64) PostResponse {
	return PostResponse{
		ID:           post.ID,
		PostTypeID:   post.PostTypeID,
		AccountID:    post.AccountID,
		Description:  post.Description.String,
		Lng:          post.Lng,
		Lat:          post.Lat,
		CreatedAt:    post.CreatedAt.Time,
		Images:       imgs,
		Account:      account,
		ReactState:   reactState,
		TotalLike:    totalLike,
		TotalComment: totalComment,
	}
}

func GetPostRes(post db.GetPostRow, account db.GetAccountByIdRow, imgs []db.PostImage, reactState db.ReactPost, totalLike, total_comment int64) PostResponse {
	return PostResponse{
		ID:           post.ID,
		PostTypeID:   post.PostTypeID,
		AccountID:    post.AccountID,
		Description:  post.Description.String,
		Lng:          post.Lng,
		Lat:          post.Lat,
		CreatedAt:    post.CreatedAt.Time,
		Images:       imgs,
		Account:      account,
		ReactState:   reactState,
		TotalLike:    totalLike,
		TotalComment: total_comment,
	}
}

// Util
func ConvertDescription(description string) pgtype.Text {
	if description == "" {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: description, Valid: true}
}
