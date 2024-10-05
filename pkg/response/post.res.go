package response

import "github.com/LeMinh0706/SocialMediaFood-Backend/db"

func PostRes(post db.Post, images []ImageResponse, user db.GetUserByIdRow, date_created int64) PostResponse {
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
