package report

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/account"
)

type ReportService struct {
	queries *db.Queries
	acc     account.IAccountService
}

// ReportPost implements IReportService.
func (r *ReportService) CreateReportPost(ctx context.Context, username string, account_id int64, post_id int64, issue_id []int32) ([]ReportPostResponse, error) {
	var reports []ReportPostResponse
	if _, err := r.acc.GetAccountAction(ctx, account_id, username); err != nil {
		return reports, err
	}
	for _, element := range issue_id {
		report, err := r.queries.CreateReport(ctx, db.CreateReportParams{
			AccountID: account_id,
			PostID:    post_id,
			IssueID:   element,
		})
		if err != nil {
			return reports, err
		}
		response := ReportPostRes(report)
		reports = append(reports, response)
	}
	return reports, nil
}

// GetListIssue implements IReportService.
func (r *ReportService) GetListIssue(ctx context.Context) []db.IssuePost {
	list, _ := r.queries.GetListIssue(ctx)
	return list
}

// GetYourReport implements IReportService.
func (r *ReportService) GetYourReport(ctx context.Context, username string, account_id int64, post_id int64) ([]ReportResponse, error) {
	var res []ReportResponse
	_, err := r.acc.GetAccountAction(ctx, account_id, username)
	if err != nil {
		return res, err
	}
	list, _ := r.queries.GetYourReport(ctx, db.GetYourReportParams{
		AccountID: account_id,
		PostID:    post_id,
	})
	for _, element := range list {
		issue, _ := r.queries.GetIssue(ctx, element.IssueID)
		res = append(res, ReportResponse{
			Id:    element.ID,
			Issue: issue,
		})
	}
	return res, nil
}

func NewReportService(queries *db.Queries, a account.IAccountService) IReportService {
	return &ReportService{
		queries: queries,
		acc:     a,
	}
}
