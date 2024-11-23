package reset_password

type ResponseLink struct {
	Link string `json:"link"`
}

type ChangePasswordRequest struct {
	Token       string `json:"kamehameha" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}
