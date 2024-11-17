package report

import "github.com/LeMinh0706/SocialMediaFood-Backend/pkg/token"

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
