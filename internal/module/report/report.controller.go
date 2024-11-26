package report

import (
	"strconv"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/handler"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/middlewares"
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

// ReportPost godoc
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

// ReportPost godoc
// @Summary      Create list Issue
// @Description  Create list Issue to report
// @Tags         Reports
// @Accept       json
// @Produce      json
// @Param        request body ReportRequest true "request"
// @Security BearerAuth
// @Success      200  {object}  []ReportPostResponse
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /report-post [post]
func (r *ReportController) CreateReport(g *gin.Context) {
	var req ReportRequest
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 40000)
		return
	}
	list, err := r.service.CreateReportPost(g, auth.UserId, req.AccountID, req.PostID, req.IssueID)
	if err != nil {
		handler.CheckPostStringError(g, err)
		return
	}
	response.SuccessResponse(g, 201, list)
}

// ReportPost godoc
// @Summary      Get your reports
// @Description  Get your reports on this post
// @Tags         Reports
// @Accept       json
// @Produce      json
// @Param        post_id query int true "PostID"
// @Param        account_id query int true "AccountID"
// @Security BearerAuth
// @Success      200  {object}  []db.IssuePost
// @Failure      500  {object}  response.ErrSwaggerJson
// @Router       /report-post/your-report [get]
func (r *ReportController) GetYourReport(g *gin.Context) {
	postStr := g.Query("post_id")
	accountStr := g.Query("account_id")
	post_id, err := strconv.ParseInt(postStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrBadRequestId)
		return
	}
	account_id, err := strconv.ParseInt(accountStr, 10, 64)
	if err != nil {
		response.ErrorResponse(g, response.ErrAccountID)
		return
	}
	auth := g.MustGet(middlewares.AuthorizationPayloadKey).(*token.Payload)
	list, err := r.service.GetYourReport(g, auth.UserId, account_id, post_id)
	if err != nil {
		handler.CheckPostStringError(g, err)
		return
	}
	response.SuccessResponse(g, 200, list)
}
