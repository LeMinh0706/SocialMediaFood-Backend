package notification

import (
	"context"
	"fmt"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/account"
	"github.com/jackc/pgx/v5/pgtype"
)

type NotificationService struct {
	queries *db.Queries
	acc     account.IAccountService
}

// CreateActionNotification implements INotificationService.
func (n *NotificationService) CreateActionNotification(ctx context.Context, account_id int64, user_action_id int64, type_id int32) error {
	panic("unimplemented")
}

// CreatePostNotification implements INotificationService.
func (n *NotificationService) CreatePostNotification(ctx context.Context, account_id int64, user_action_id int64, post_id int64, type_id int32) error {

	name, _ := n.acc.GetAccountById(ctx, user_action_id)
	post, _ := n.queries.GetPost(ctx, post_id)
	des := []rune(post.Description.String)

	summary := string(des)
	if len(des) > 10 {
		summary = string(des[:10])
	}
	var message string
	switch type_id {
	case 1:
		message = fmt.Sprintf("%v đã bình luận bài viết của bạn: %v...", name.Fullname, summary)
	case 2:
		message = fmt.Sprintf("%v đã thích bài viết của bạn: %v...", name.Fullname, summary)
	}

	_, err := n.queries.CreatePostNoti(ctx, db.CreatePostNotiParams{
		Message:   message,
		AccountID: account_id,
		TypeID:    type_id,
		PostID: pgtype.Int8{
			Int64: post_id,
			Valid: true,
		},
		UserActionID: user_action_id,
	})
	if err != nil {
		return err
	}
	return nil
}

// DeleteNoti implements INotificationService.
func (n *NotificationService) DeleteNoti(ctx context.Context, user_id int64, id int64) error {
	noti, err := n.queries.GetNotification(ctx, id)
	if err != nil {
		return err
	}
	_, err = n.acc.GetAccountAction(ctx, noti.AccountID, user_id)
	if err != nil {
		return err
	}
	err = n.queries.DeleteNoti(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// GetListNotification implements INotificationService.
func (n *NotificationService) GetListNotification(ctx context.Context, user_id int64, account_id int64, page int32, pageSize int32) ([]NotificationResponse, error) {
	var res []NotificationResponse
	_, err := n.acc.GetAccountAction(ctx, account_id, user_id)
	if err != nil {
		return res, err
	}
	list, err := n.queries.GetListNoti(ctx, db.GetListNotiParams{
		AccountID: account_id,
		Limit:     pageSize,
		Offset:    (page - 1) * pageSize,
	})
	if err != nil {
		return res, err
	}
	for _, element := range list {
		noti, err := n.GetNotification(ctx, element.ID, element.UserActionID)
		if err != nil {
			return []NotificationResponse{}, err
		}
		res = append(res, noti)
	}
	return res, nil
}

// GetNotification implements INotificationService.
func (n *NotificationService) GetNotification(ctx context.Context, id int64, user_action_id int64) (NotificationResponse, error) {
	var res NotificationResponse
	other, _ := n.acc.GetAccountById(ctx, user_action_id)
	noti, err := n.queries.GetNotification(ctx, id)
	if err != nil {
		return res, err
	}
	res = NotiRes(noti, other)
	return res, nil
}

// IsSeen implements INotificationService.
func (n *NotificationService) IsSeen(ctx context.Context, user_id int64, id int64) error {
	noti, err := n.queries.GetNotification(ctx, id)
	if err != nil {
		return err
	}
	_, err = n.acc.GetAccountAction(ctx, noti.AccountID, user_id)
	if err != nil {
		return err
	}

	err = n.queries.UpdateSeen(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// IsSeenAll implements INotificationService.
func (n *NotificationService) IsSeenAll(ctx context.Context, user_id int64, account_id int64) error {
	_, err := n.acc.GetAccountAction(ctx, account_id, user_id)
	if err != nil {
		return err
	}
	err = n.queries.UpdateSeenAll(ctx, account_id)
	if err != nil {
		return err
	}
	return nil
}

func NewNotificationService(q *db.Queries, a account.IAccountService) INotificationService {
	return &NotificationService{
		queries: q,
		acc:     a,
	}
}
