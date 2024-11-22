package notification

import (
	"strconv"
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type NotificationResponse struct {
	ID         int64                `json:"id"`
	Message    string               `json:"message"`
	AccountID  int64                `json:"account_id"`
	TypeID     int32                `json:"type_id"`
	PostID     int64                `json:"post_id"`
	UserAction db.GetAccountByIdRow `json:"user_action"`
	InvoiceID  int64                `json:"invoice_id"`
	IsSeen     bool                 `json:"is_seen"`
	CreatedAt  time.Time            `json:"created_at"`
}

func NotiRes(noti db.Notification, other db.GetAccountByIdRow) NotificationResponse {
	return NotificationResponse{
		ID:         noti.ID,
		Message:    noti.Message,
		AccountID:  noti.AccountID,
		TypeID:     noti.TypeID,
		PostID:     noti.PostID.Int64,
		UserAction: other,
		InvoiceID:  noti.InvoiceID.Int64,
		IsSeen:     noti.IsSeen,
		CreatedAt:  noti.CreatedAt.Time,
	}
}

func CheckQuery(g *gin.Context, pageStr, pageSizeStr string) (int32, int32) {
	page, err := strconv.ParseInt(pageStr, 10, 32)
	if err != nil {
		response.ErrorResponse(g, 40001)
		return 0, 0
	}
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 32)
	if err != nil {
		response.ErrorResponse(g, 40002)
		return 0, 0
	}
	return int32(page), int32(pageSize)
}
