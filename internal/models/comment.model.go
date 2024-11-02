package models

import (
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type CommentRequest struct {
	AccountID   int64  `json:"account_id" binding:"required"`
	PostTopID   int64  `json:"post_top_id" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CommentResponse struct {
	ID          int64          `json:"id"`
	Account     AccountForPost `json:"account"`
	PostTopID   int64          `json:"post_top_id"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
}

func CommentRes(acc AccountForPost, comment db.GetCommentRow) CommentResponse {
	return CommentResponse{
		ID:          comment.ID,
		Account:     acc,
		PostTopID:   comment.PostTopID.Int64,
		Description: comment.Description.String,
		CreatedAt:   comment.CreatedAt.Time,
	}
}

func CommentReq(req CommentRequest) db.CreateCommentParams {
	return db.CreateCommentParams{
		AccountID:   req.AccountID,
		PostTopID:   pgtype.Int8{Int64: req.PostTopID, Valid: true},
		Description: pgtype.Text{String: req.Description, Valid: true},
	}
}
