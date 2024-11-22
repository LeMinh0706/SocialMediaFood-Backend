package reset_password

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type ResponseLink struct {
	Link string `json:"link"`
}

type ChangePasswordRequest struct {
	Token       string `json:"kamehameha" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

func ResetPasswordErr(g *gin.Context, err error) {
	switch err.Error() {
	case "not found":
		response.ErrorResponse(g, response.ErrEmailNotExists)
		return
	case "intime":
		response.ErrorResponse(g, response.YouHaveRequest)
		return
	case "request used":
		response.ErrorResponse(g, response.PasswordHaveChange)
		return
	}
	response.ErrorNonKnow(g, 500, err.Error())
}
