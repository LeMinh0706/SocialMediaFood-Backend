package report

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type IReportService interface {
	GetListIssue(ctx context.Context) []db.IssuePost
	GetYourReport(ctx context.Context, username string, account_id, post_id int64) ([]ReportResponse, error)
	CreateReportPost(ctx context.Context, username string, account_id, post_id int64, issue_id []int32) ([]ReportPostResponse, error)
}
