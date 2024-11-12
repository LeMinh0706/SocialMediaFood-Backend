package notification

import (
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/account"
)

type NotificationResponse struct {
	ID         int64                   `json:"id"`
	Message    string                  `json:"message"`
	Account    account.AccountResponse `json:"account"`
	TypeID     int32                   `json:"type_id"`
	PostID     int64                   `json:"post_id"`
	UserAction account.AccountResponse `json:"user_action"`
	InvoiceID  int64                   `json:"invoice_id"`
	IsSeen     bool                    `json:"is_seen"`
	CreatedAt  time.Time               `json:"created_at"`
}
