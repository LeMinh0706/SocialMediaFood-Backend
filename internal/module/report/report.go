package report

import (
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type ReportRequest struct {
	PostID    int64   `json:"post_id" binding:"required"`
	AccountID int64   `json:"account_id" binding:"required"`
	IssueID   []int32 `json:"list_issue_id"`
}

type ReportResponse struct {
	Id    int64        `json:"id"`
	Issue db.IssuePost `json:"issue"`
}

type ReportPostResponse struct {
	ID        int64     `json:"id"`
	AccountID int64     `json:"account_id"`
	IssueID   int32     `json:"issue_id"`
	PostID    int64     `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}

func ReportPostRes(report db.ReportPost) ReportPostResponse {
	return ReportPostResponse{
		ID:        report.ID,
		AccountID: report.AccountID,
		IssueID:   report.IssueID,
		PostID:    report.PostID,
		CreatedAt: report.CreatedAt.Time,
	}
}
