package response

import "github.com/LeMinh0706/SocialMediaFood-Backend/db"

type CommentRequest struct {
	UserID      int64  `json:"user_id" binding:"required"`
	PostTopID   int64  `json:"post_top_id" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateCommentRequest struct {
	ID          int64  `json:"id" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CommentResponse struct {
	ID             int64             `json:"id"`
	PostTopID      int64             `json:"post_top_id"`
	Description    string            `json:"description"`
	DateCreatePost int64             `json:"date_create_post"`
	User           db.GetUserByIdRow `json:"user"`
}

func CommentRes(comment db.Post, user db.GetUserByIdRow) CommentResponse {
	return CommentResponse{
		ID:             comment.ID,
		PostTopID:      comment.PostTopID.Int64,
		Description:    comment.Description.String,
		DateCreatePost: comment.DateCreatePost,
		User:           user,
	}
}
