package user

import (
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Login struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=6,max=16" example:"HiroPhent"`
	Password string `json:"password" binding:"required,min=8" example:"kocanpass"`
}

type LoginResponse struct {
	Token string `json:"access_token"`
}

type RegisterResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func ValidateRegister(g *gin.Context, err error) {
	if validate, ok := err.(validator.ValidationErrors); ok {
		for _, vali := range validate {
			switch vali.Tag() {
			case "min":
				if vali.Field() == "Username" {
					response.ErrorResponse(g, 40008)
					return
				} else if vali.Field() == "Gender" {
					response.ErrorResponse(g, 40007)
					return
				} else if vali.Field() == "Password" {
					response.ErrorResponse(g, 40009)
					return
				} else if vali.Field() == "Fullname" {
					response.ErrorResponse(g, 40010)
					return
				}
			case "max":
				if vali.Field() == "Username" {
					response.ErrorResponse(g, 40008)
					return
				} else if vali.Field() == "Gender" {
					response.ErrorResponse(g, 40007)
					return
				}
			case "required":
				if vali.Field() == "Username" {
					response.ErrorResponse(g, 40008)
					return
				} else if vali.Field() == "Gender" {
					response.ErrorResponse(g, 40007)
					return
				} else if vali.Field() == "Password" {
					response.ErrorResponse(g, 40009)
					return
				} else if vali.Field() == "Fullname" {
					response.ErrorResponse(g, 40010)
					return
				}
			}
		}
		response.ErrorNonKnow(g, 400, err.Error())
		return
	}
}
