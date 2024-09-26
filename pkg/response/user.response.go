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

type UserForPost struct {
	ID       int64  `json:"id"`
	Fullname string `json:"fullname"`
	RoleID   int32  `json:"role_id"`
}

type UserResponse struct {
	ID                int64  `json:"id"`
	Fullname          string `json:"fullname"`
	Gender            int32  `json:"gender"`
	RoleID            int32  `json:"role_id"`
	DateCreateAccount int64  `json:"date_create_account"`
}
