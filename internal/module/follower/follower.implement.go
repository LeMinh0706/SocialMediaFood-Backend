package follower

import (
	"context"
	"fmt"

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
	to_acc, err := f.account.GetAccountById(ctx, to.FromFollow)
	if err != nil {
		return from_res, to_res, err
	}
	from_follow := AccountFollowResponse{Account: from_acc, Status: from.Status}
	to_follow := AccountFollowResponse{Account: to_acc, Status: to.Status}
	return from_follow, to_follow, err
}

// FollowRequest implements IFollowerService.
func (f *FollowerService) FollowRequest(ctx context.Context, username string, from_id int64, to_id int64) (FollowResponse, error) {
	var res FollowResponse
	if from_id == to_id {
		return res, fmt.Errorf("can't follow yourself")
	}
	_, err := f.account.GetAccountAction(ctx, from_id, username)
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
func (f *FollowerService) GetYourFollower(ctx context.Context, page int32, page_size int32, from_id int64) (ListFollow, error) {
	var res ListFollow
	list, _ := f.queries.GetYourFollower(ctx, db.GetYourFollowerParams{
		FromFollow: from_id,
		Limit:      page_size,
		Offset:     (page - 1) * page_size,
	})
	accounts := make([]db.GetAccountByIdRow, 0)
	for _, element := range list {
		follower, _ := f.account.GetAccountById(ctx, element)
		accounts = append(accounts, follower)
	}
	total, _ := f.queries.CountFollower(ctx, from_id)
	res = ListFollow{Account: accounts, Total: total}
	return res, nil
}

// GetYourRequest implements IFollowerService.
func (f *FollowerService) GetYourRequest(ctx context.Context, page int32, page_size int32, from_id int64) (ListFollow, error) {
	var res ListFollow
	list, _ := f.queries.GetYourRequest(ctx, db.GetYourRequestParams{
		FromFollow: from_id,
		Limit:      page_size,
		Offset:     (page - 1) * page_size,
	})
	accounts := make([]db.GetAccountByIdRow, 0)
	for _, element := range list {
		follower, _ := f.account.GetAccountById(ctx, element)
		accounts = append(accounts, follower)
	}
	total, _ := f.queries.CountRequest(ctx, from_id)
	res = ListFollow{Account: accounts, Total: total}
	return res, nil
}

// GetYourFriend implements IFollowerService.
func (f *FollowerService) GetYourFriend(ctx context.Context, page int32, page_size int32, from_id int64) (ListFollow, error) {
	var res ListFollow
	list, _ := f.queries.GetYourFriend(ctx, db.GetYourFriendParams{
		FromFollow: from_id,
		Limit:      page_size,
		Offset:     (page - 1) * page_size,
	})
	accounts := make([]db.GetAccountByIdRow, 0)
	for _, element := range list {
		follower, _ := f.account.GetAccountById(ctx, element)
		accounts = append(accounts, follower)
	}
	total, _ := f.queries.CountFriend(ctx, from_id)
	res = ListFollow{Account: accounts, Total: total}
	return res, nil
}

// GetFollowType implements IFollowerService.
func (f *FollowerService) GetFollowType(ctx context.Context, status string, page int32, page_size int32, from_id int64) (ListFollow, error) {
	switch status {
	case "accept":
		follow, err := f.GetYourFollower(ctx, page, page_size, from_id)
		if err != nil {
			return ListFollow{}, err
		}
		return follow, nil
	case "request":
		follow, err := f.GetYourRequest(ctx, page, page_size, from_id)
		if err != nil {
			return ListFollow{}, err
		}
		return follow, nil
	case "friend":
		follow, err := f.GetYourFriend(ctx, page, page_size, from_id)
		if err != nil {
			return ListFollow{}, err
		}
		return follow, nil
	default:
		return ListFollow{}, fmt.Errorf("wrong status input")
	}
}

// UpdateStatus implements IFollowerService.
func (f *FollowerService) UpdateStatus(ctx context.Context, username string, from_id int64, to_id int64) error {
	_, err := f.account.GetAccountAction(ctx, from_id, username)
	if err != nil {
		return err
	}
	follow, err := f.GetRequestStatus(ctx, db.GetFollowStatusParams{
		FromFollow: from_id,
		ToFollow:   to_id,
	})
	if err != nil {
		return err
	}
	if follow.Status != "accept" {
		return fmt.Errorf("wait reply")
	}
	err = f.queries.UpdateFriend(ctx, db.UpdateFriendParams{FromFollow: from_id, ToFollow: to_id})
	if err != nil {
		return err
	}
	return nil
}

// UnFollow implements IFollowerService.
func (f *FollowerService) UnFollow(ctx context.Context, username string, from_id int64, to_id int64) error {
	_, err := f.account.GetAccountAction(ctx, from_id, username)
	if err != nil {
		return err
	}
	return f.queries.DeleteFollow(ctx, db.DeleteFollowParams{
		FromFollow: from_id,
		ToFollow:   to_id,
	})
}

func NewFollowerService(queries *db.Queries, acc account.IAccountService) IFollowerService {
	return &FollowerService{
		queries: queries,
		account: acc,
	}
}
