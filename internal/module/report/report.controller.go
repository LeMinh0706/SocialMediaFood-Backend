package report

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

type ReportController struct {
	service IReportService
	token   token.Maker
}

func NewReportController(service IReportService, token token.Maker) *ReportController {
	return &ReportController{
		service: service,
		token:   token,
	}
}

// Post godoc
// @Summary      Get list Issue
// @Description  Get list Issue to report
// @Tags         Reports
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Success      200  {object}  []db.IssuePost
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /report-post/issue [get]
func (r *ReportController) GetListIssue(g *gin.Context) {
	list := r.service.GetListIssue(g)
	response.SuccessResponse(g, 200, list)
}
