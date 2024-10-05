package response

type RegisterResponse struct {
	ID                int64  `json:"id"`
	Email             string `json:"email"`
	Username          string `json:"username"`
	Fullname          string `json:"fullname"`
	Gender            int32  `json:"gender"`
	RoleID            int32  `json:"role_id"`
	DateCreateAccount int64  `json:"date_create_account"`
}

type RequestResponse struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6,max=18"`
	Gender   int32  `json:"gender" binding:"required,min=0,max=1"`
}

type UserForPost struct {
	ID        int64  `json:"id"`
	Fullname  string `json:"fullname"`
	RoleID    int32  `json:"role_id"`
	UrlAvatar string `json:"url_avatar"`
}

type UserResponse struct {
	ID                int64  `json:"id"`
	Fullname          string `json:"fullname"`
	Gender            int32  `json:"gender"`
	RoleID            int32  `json:"role_id"`
	UrlAvatar         string `json:"url_avatar"`
	UrlBackground     string `json:"url_background"`
	DateCreateAccount int64  `json:"date_create_account"`
}

type RequestLogin struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6,max=18"`
}

type LoginResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user"`
}
