package react

import (
	"context"
	"fmt"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/account"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/notification"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/post"
	"github.com/jackc/pgx/v5"
)

type ReactService struct {
	queries *db.Queries
	account account.IAccountService
	post    post.IPostService
	noti    notification.INotificationService
}

// Backup implements IReactService.
func (r *ReactService) Backup(ctx context.Context) {
	panic("unimplemented")
}

// ChangeReactState implements IReactService.
func (r *ReactService) ChangeReactState(ctx context.Context, user_id int64, account_id int64, post_id int64, state int32) (db.ReactPost, error) {
	var res db.ReactPost
	_, err := r.account.GetAccountAction(ctx, account_id, user_id)
	if err != nil {
		return res, err
	}
	_, err = r.GetReactPost(ctx, account_id, post_id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return res, fmt.Errorf("err like")
		}
		return res, err
	}
	update, err := r.queries.UpdateState(ctx, db.UpdateStateParams{
		PostID:    post_id,
		AccountID: account_id,
		State:     state,
	})
	if err != nil {
		return res, err
	}
	return update, nil
}

// CreateReact implements IReactService.
func (r *ReactService) CreateReact(ctx context.Context, user_id int64, account_id int64, post_id int64) (db.ReactPost, error) {
	var res db.ReactPost
	_, err := r.account.GetAccountAction(ctx, account_id, user_id)
	if err != nil {
		return res, err
	}
	p, err := r.post.GetPost(ctx, account_id, post_id)
	if err != nil {
		return res, err
	}
	react, err := r.queries.CreateReact(ctx, db.CreateReactParams{
		AccountID: account_id,
		PostID:    post_id,
		State:     1,
	})
	if err != nil {
		return res, err
	}
	if p.AccountID != account_id {
		r.noti.CreatePostNotification(ctx, p.AccountID, account_id, post_id, 2)
	}

	return react, err
}

// GetListReactPost implements IReactService.
func (r *ReactService) GetListReactPost(ctx context.Context, page int32, pageSize int32, post_id int64) (ListReactResponse, error) {
	var res ListReactResponse
	var a []ReactResponse
	list, err := r.queries.GetListReact(ctx, db.GetListReactParams{
		PostID: post_id,
		Limit:  pageSize,
		Offset: (page - 1) * pageSize,
	})
	if err != nil {
		return res, err
	}
	for _, element := range list {
		react, err := r.GetReactPost(ctx, element, post_id)
		if err != nil {
			return ListReactResponse{}, err
		}
		a = append(a, react)
	}
	total, _ := r.queries.CountReactPost(ctx, post_id)
	res = ListReactResponse{React: a, Total: total}
	return res, nil
}

// GetReactPost implements IReactService.
func (r *ReactService) GetReactPost(ctx context.Context, account_id int64, post_id int64) (ReactResponse, error) {
	var res ReactResponse
	account, _ := r.account.GetAccountById(ctx, account_id)
	react, err := r.queries.GetReact(ctx, db.GetReactParams{
		AccountID: account_id,
		PostID:    post_id,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return ReactResponse{
				PostID:  post_id,
				Account: account,
				State:   0,
			}, nil
		}
		return res, err
	}
	res = ReactResponse{
		PostID:  post_id,
		Account: account,
		State:   react.State,
	}
	return res, nil
}

// UnReaction implements IReactService.
func (r *ReactService) UnReaction(ctx context.Context, user_id int64, account_id int64, post_id int64) error {
	_, err := r.account.GetAccountAction(ctx, account_id, user_id)
	if err != nil {
		return err
	}
	_, err = r.GetReactPost(ctx, account_id, post_id)
	if err != nil {
		return err
	}
	err = r.queries.DeleteReact(ctx, db.DeleteReactParams{
		PostID:    post_id,
		AccountID: account_id,
	})
	if err != nil {
		return fmt.Errorf("err like")
	}
	return nil
}

func NewReactService(queries *db.Queries, acc account.IAccountService, post post.IPostService, ns notification.INotificationService) IReactService {
	return &ReactService{
		queries: queries,
		account: acc,
		post:    post,
		noti:    ns,
	}
}
