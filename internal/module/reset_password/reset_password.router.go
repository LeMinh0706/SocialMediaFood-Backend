package reset_password

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"github.com/gin-gonic/gin"
)

func NewResetPasswordRouter(r *gin.Engine, group *gin.RouterGroup, service IResetPasswordService, config util.Config, token token.Maker) {
	rc := NewResetPasswordController(service, config, token)
	reset := r.Group(group.BasePath() + "/reset-password")
	{
		reset.POST("", rc.ResetPasswordRequest)
	}
}
