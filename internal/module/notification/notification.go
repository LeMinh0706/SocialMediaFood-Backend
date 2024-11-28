package notification

import (
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
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
	IsDelete   bool                 `json:"is_delete"`
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
		IsDelete:   noti.IsDelete,
	}
}
