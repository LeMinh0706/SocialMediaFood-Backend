package notification

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/account"
)

type NotificationService struct {
	queries *db.Queries
	acc     account.IAccountService
}

// CreateActionNotification implements INotificationService.
func (n *NotificationService) CreateActionNotification(ctx context.Context, arg db.CreateActionNotiParams) error {
	_, err := n.queries.CreateActionNoti(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}

// CreatePostNotification implements INotificationService.
func (n *NotificationService) CreatePostNotification(ctx context.Context, arg db.CreatePostNotiParams) error {
	_, err := n.queries.CreatePostNoti(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}

// DeleteNoti implements INotificationService.
func (n *NotificationService) DeleteNoti(ctx context.Context, user_id int64, id int64) error {
	panic("unimplemented")
}

// GetListNotification implements INotificationService.
func (n *NotificationService) GetListNotification(ctx context.Context, user_id int64, account_id int64) ([]NotificationResponse, error) {
	panic("unimplemented")
}

// GetNotification implements INotificationService.
func (n *NotificationService) GetNotification(ctx context.Context, id int64) (NotificationResponse, error) {
	panic("unimplemented")
}

// IsSeen implements INotificationService.
func (n *NotificationService) IsSeen(ctx context.Context, user_id int64, id int64) error {
	panic("unimplemented")
}

// IsSeenAll implements INotificationService.
func (n *NotificationService) IsSeenAll(ctx context.Context, user_id int64, account_id int64) error {
	panic("unimplemented")
}

func NewNotificationService(q *db.Queries, acccountService account.IAccountService) INotificationService {
	return &NotificationService{
		queries: q,
		acc:     acccountService,
	}
}
