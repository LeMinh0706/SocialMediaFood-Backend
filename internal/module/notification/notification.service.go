package notification

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type INotificationService interface {
	CreatePostNotification(ctx context.Context, arg db.CreatePostNotiParams) error
	CreateActionNotification(ctx context.Context, arg db.CreateActionNotiParams) error
	GetNotification(ctx context.Context, id int64) (NotificationResponse, error)
	GetListNotification(ctx context.Context, user_id, account_id int64) ([]NotificationResponse, error)
	IsSeen(ctx context.Context, user_id int64, id int64) error
	IsSeenAll(ctx context.Context, user_id, account_id int64) error
	DeleteNoti(ctx context.Context, user_id, id int64) error
}
