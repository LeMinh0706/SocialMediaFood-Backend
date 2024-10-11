package response

type ReactPostResponse struct {
	PostID int64               `json:"post_id"`
	Users  []UserReactResponse `json:"users"`
	Total  int64               `json:"Total"`
}

type UserReactResponse struct {
	UserID int64 `json:"user_id"`
}

func ReactPostRes(post_id, total int64) ReactPostResponse {
	return ReactPostResponse{
		PostID: post_id,
		Total:  total,
	}
}
