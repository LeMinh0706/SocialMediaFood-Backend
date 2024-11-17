package report

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewReportRouter(r *gin.Engine, group *gin.RouterGroup, service IReportService, token token.Maker) {
	// rc := NewReportController(service, token)
	// reportGroup := r.Group("/report-post")
	// auth := reportGroup.Group("").Use(middlewares.AuthorizeMiddleware(token))
	// {
	// }
}
