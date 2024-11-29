package admin

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewAdminRouter(r *gin.Engine, group *gin.RouterGroup, service IAdminService, token token.Maker) {
}
