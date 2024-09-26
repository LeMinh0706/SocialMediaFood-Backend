package response

type PostResponse struct {
	ID             int64  `json:"id"`
	PostTypeID     int32  `json:"post_type_id"`
	UserID         int64  `json:"user_id"`
	Description    string `json:"description"`
	DateCreatePost int64  `json:"date_create_post"`
}
