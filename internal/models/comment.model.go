package models

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type CommentRequest struct {
	AccountID   int64  `json:"account_id" binding:"required"`
	PostTopID   int64  `json:"post_top_id" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func CommentReq(req CommentRequest) db.CreateCommentParams {
	return db.CreateCommentParams{
		AccountID:   req.AccountID,
		PostTopID:   pgtype.Int8{Int64: req.PostTopID, Valid: true},
		Description: pgtype.Text{String: req.Description, Valid: true},
	}
}
