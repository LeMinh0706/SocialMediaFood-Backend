package comment

import (
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type CommentResponse struct {
	ID          int64                `json:"id"`
	PostTypeID  int32                `json:"post_type_id"`
	AccountID   int64                `json:"account_id"`
	PostTopID   int64                `json:"post_top_id"`
	Description string               `json:"description"`
	CreatedAt   time.Time            `json:"created_at"`
	Image       db.PostImage         `json:"image"`
	Account     db.GetAccountByIdRow `json:"account"`
}

func CommentRes(comment db.CreateCommentRow, image db.PostImage, account db.GetAccountByIdRow) CommentResponse {
	return CommentResponse{
		ID:          comment.ID,
		PostTypeID:  comment.PostTypeID,
		AccountID:   comment.AccountID,
		PostTopID:   comment.PostTopID.Int64,
		Description: comment.Description.String,
		CreatedAt:   comment.CreatedAt.Time,
		Image:       image,
		Account:     account,
	}
}
