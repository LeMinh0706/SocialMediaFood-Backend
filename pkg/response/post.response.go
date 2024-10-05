package response

import "github.com/LeMinh0706/SocialMediaFood-Backend/db"

// type PostRequest struct {
// 	Description string   `json:"description"`
// 	UserId      int64    `json:"user_id" binding:"required"`
// 	Images      []string `json:"image_url" binding:"required"`
// }

type PostResponse struct {
	ID             int64             `json:"id"`
	PostTypeID     int32             `json:"post_type_id"`
	UserID         int64             `json:"user_id"`
	Description    string            `json:"description"`
	Images         []db.PostImage    `json:"images"`
	User           db.GetUserByIdRow `json:"user"`
	DateCreatePost int64             `json:"date_create_post"`
}

func PostRes(post db.Post, images []db.PostImage, user db.GetUserByIdRow, date_created int64) PostResponse {
	return PostResponse{
		ID:             post.ID,
		PostTypeID:     post.PostTypeID,
		UserID:         user.ID,
		Description:    post.Description.String,
		Images:         images,
		User:           user,
		DateCreatePost: date_created,
	}
}
