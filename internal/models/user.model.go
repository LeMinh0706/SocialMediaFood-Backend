package models

import (
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type RegisterResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginRequest struct {
	Username string `json:"username" example:"HiroPhent"`
	Password string `json:"password" example:"kocanpass"`
}

type LoginResponse struct {
	Token string `json:"access_token"`
}

func RegisterRes(user db.RegisterRow) RegisterResponse {
	return RegisterResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email.String,
		CreatedAt: user.CreatedAt.Time,
	}
}

// func LoginRes(user db.LoginRow, token string) LoginResponse {
// 	return LoginResponse{
// 		Token: token,
// 	}
// }
