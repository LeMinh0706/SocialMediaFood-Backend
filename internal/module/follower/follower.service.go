package follower

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/account"
	"github.com/jackc/pgx/v5"
)

type FollowerService struct {
	queries *db.Queries
	account account.IAccountService
}

// Backup implements IFollowerService.
func (f *FollowerService) Backup(ctx context.Context) {
	panic("unimplemented")
}

// UtilCreateFollow implements IFollowerService.
func (f *FollowerService) UtilCreateFollow(ctx context.Context, from db.Follower, to db.Follower) (AccountFollowResponse, AccountFollowResponse, error) {
	var from_res AccountFollowResponse
	var to_res AccountFollowResponse
	from_acc, err := f.account.GetAccountById(ctx, from.FromFollow)
	if err != nil {
		return from_res, to_res, err
	}
	to_acc, err := f.account.GetAccountById(ctx, to.ToFollow)
	if err != nil {
		return from_res, to_res, err
	}
	from_follow := AccountFollowResponse{Account: from_acc, Status: from.Status}
	to_follow := AccountFollowResponse{Account: to_acc, Status: to.Status}
	return from_follow, to_follow, err
}

// FollowRequest implements IFollowerService.
func (f *FollowerService) FollowRequest(ctx context.Context, user_id int64, from_id int64, to_id int64) (FollowResponse, error) {
	var res FollowResponse
	_, err := f.account.GetAccountAction(ctx, from_id, user_id)
	if err != nil {
		return res, err
	}
	from, err := f.queries.CreateFollow(ctx, db.CreateFollowParams{
		FromFollow: from_id,
		ToFollow:   to_id,
		Status:     "request",
	})
	if err != nil {
		return res, err
	}
	to, err := f.queries.CreateFollow(ctx, db.CreateFollowParams{
		FromFollow: to_id,
		ToFollow:   from_id,
		Status:     "accept",
	})
	if err != nil {
		return res, err
	}
	from_follow, to_follow, err := f.UtilCreateFollow(ctx, from, to)
	if err != nil {
		return res, err
	}
	res = FollowResponse{
		FromFollow: from_follow,
		ToFollow:   to_follow,
	}
	return res, nil
}

// GetRequestStatus implements IFollowerService.
func (f *FollowerService) GetRequestStatus(ctx context.Context, arg db.GetFollowStatusParams) (db.GetFollowStatusRow, error) {
	var res db.GetFollowStatusRow
	_, err := f.account.GetAccountById(ctx, arg.FromFollow)
	if err != nil {
		return res, err
	}
	_, err = f.account.GetAccountById(ctx, arg.ToFollow)
	if err != nil {
		return res, err
	}
	status, err := f.queries.GetFollowStatus(ctx, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			return db.GetFollowStatusRow{
				FromFollow: arg.FromFollow,
				ToFollow:   arg.ToFollow,
				Status:     "no status",
			}, nil
		}
		return res, err
	}
	return status, nil
}

// GetYourFollower implements IFollowerService.
func (f *FollowerService) GetYourFollower(ctx context.Context, from_id int64, to_id int64) ([]account.AccountResponse, error) {
	panic("unimplemented")
}

// GetYourFriend implements IFollowerService.
func (f *FollowerService) GetYourFriend(ctx context.Context, from_id int64, to_id int64) ([]account.AccountResponse, error) {
	panic("unimplemented")
}

// GetYourRequest implements IFollowerService.
func (f *FollowerService) GetYourRequest(ctx context.Context, from_id int64, to_id int64) ([]account.AccountResponse, error) {
	panic("unimplemented")
}

// UnFollow implements IFollowerService.
func (f *FollowerService) UnFollow(ctx context.Context, from_id int64, to_id int64) error {
	panic("unimplemented")
}

// UpdateStatus implements IFollowerService.
func (f *FollowerService) UpdateStatus(ctx context.Context, user_id int64, from_id int64, to_id int64) (FollowResponse, error) {
	panic("unimplemented")
}

func NewFollowerService(queries *db.Queries, acc account.IAccountService) IFollowerService {
	return &FollowerService{
		queries: queries,
		account: acc,
	}
}
