package notification

import (
	"context"
)

type INotificationService interface {
	CreatePostNotification(ctx context.Context, account_id, user_action_id, post_id int64, type_id int32) error
	CreateActionNotification(ctx context.Context, account_id, user_action_id int64, type_id int32) error
	GetNotification(ctx context.Context, id int64, user_action_id int64) (NotificationResponse, error)
	GetListNotification(ctx context.Context, username string, account_id int64, page, pageSize int32) ([]NotificationResponse, error)
	IsSeen(ctx context.Context, username string, id int64) error
	IsSeenAll(ctx context.Context, username string, account_id int64) error
	DeleteNoti(ctx context.Context, username string, id int64) error
}
