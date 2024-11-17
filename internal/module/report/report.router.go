package report

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewReportRouter(r *gin.Engine, group *gin.RouterGroup, service IReportService, token token.Maker) {
	rc := NewReportController(service, token)
	reportGroup := r.Group(group.BasePath() + "/report-post")
	auth := reportGroup.Group("").Use(middlewares.AuthorizeMiddleware(token))
	{
		auth.GET("/issue", rc.GetListIssue)
		auth.GET("/your-report", rc.GetYourReport)
		auth.POST("", rc.CreateReport)
	}

}
