package user

import (
	"time"
)

type Login struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type AccessRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type AccessResponse struct {
	AccessToken string `json:"access_token"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=6,max=16" example:"HiroPhent"`
	Password string `json:"password" binding:"required,min=8" example:"kocanpass"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	RefeshToken string `json:"refresh_token"`
}

type RegisterResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
