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

// CreateReportPost implements IReportService.
func (r *ReportService) CreateReportPost(ctx context.Context, user_id int64, account_id int64, post_id int64, issue_id []int32) ([]db.ReportPost, error) {
	var reports []db.ReportPost
	if _, err := r.acc.GetAccountAction(ctx, account_id, user_id); err != nil {
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
		reports = append(reports, report)
	}
	return reports, nil
}

// GetListIssue implements IReportService.
func (r *ReportService) GetListIssue(ctx context.Context) []db.IssuePost {
	list, _ := r.queries.GetListIssue(ctx)
	return list
}

// GetYourReport implements IReportService.
func (r *ReportService) GetYourReport(ctx context.Context, user_id int64, account_id int64, post_id int64) ([]db.GetYourReportRow, error) {
	panic("unimplemented")
}

func NewReportService(queries *db.Queries, a account.IAccountService) IReportService {
	return &ReportService{
		queries: queries,
		acc:     a,
	}
}
