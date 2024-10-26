package response

import (
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
	Fullname string `json:"fullname" binding:"required"`
	Gender   int32  `json:"gender" binding:"min=0,max=1"`
}

type RegisterResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Token     string    `json:"access_token"`
}

func RegisterRes(user db.RegisterRow) RegisterResponse {
	return RegisterResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email.String,
		CreatedAt: user.CreatedAt.Time,
	}
}

func LoginRes(user db.LoginRow, token string) LoginResponse {
	return LoginResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email.String,
		CreatedAt: user.CreatedAt.Time,
		Token:     token,
	}
}
