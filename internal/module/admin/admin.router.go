package admin

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewAdminRouter(r *gin.Engine, group *gin.RouterGroup, service IAdminService, token token.Maker) {
	ac := NewAdminController(service)
	adminGroup := r.Group(group.BasePath() + "/admin").Use(middlewares.AuthorizeMiddleware(token))
	{
		adminGroup.POST("/price", ac.AddUpgradePrice)
		adminGroup.GET("/price", ac.GetListUpgradePrice)
		adminGroup.GET("/report", ac.GetListReportPost)
		adminGroup.GET("/report/:id", ac.GetDetailReportPost)
		adminGroup.GET("/upgrade-queue", ac.GetUpgradeQueue)
		adminGroup.POST("/upgrade-queue/:id", ac.UpgradeSuccess)
		adminGroup.POST("/price-choosing/:id", ac.PriceChoosing)
		adminGroup.DELETE("/upgrade-queue/:id", ac.UpgradeReject)
		adminGroup.POST("/ban-post/:id", ac.BanPost)
		adminGroup.DELETE("/ban-post/:id", ac.RejectBan)
	}
}
