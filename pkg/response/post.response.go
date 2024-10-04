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
	Images         []ImageResponse   `json:"images"`
	User           db.GetUserByIdRow `json:"user"`
	DateCreatePost int64             `json:"date_create_post"`
}

type ImageResponse struct {
	ID       int64  `json:"id"`
	UrlImage string `json:"url_image"`
	PostId   int64  `json:"post_id"`
}
