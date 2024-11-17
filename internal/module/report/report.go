package report

import "github.com/LeMinh0706/SocialMediaFood-Backend/db"

type ReportRequest struct {
	PostID    int64   `json:"post_id" binding:"required"`
	AccountID int64   `json:"account_id" binding:"required"`
	IssueID   []int32 `json:"list_issue_id"`
}

type ReportResponse struct {
	Id    int64        `json:"id"`
	Issue db.IssuePost `json:"issue"`
}
