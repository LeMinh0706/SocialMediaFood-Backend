package follower

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/account"
)

type IFollowerService interface {
	Backup(ctx context.Context)
	FollowRequest(ctx context.Context, user_id, from_id, to_id int64) (FollowResponse, error)
	UtilCreateFollow(ctx context.Context, from, to db.Follower) (AccountFollowResponse, AccountFollowResponse, error)
	GetRequestStatus(ctx context.Context, arg db.GetFollowStatusParams) (db.GetFollowStatusRow, error)
	GetYourFollower(ctx context.Context, from_id, to_id int64) ([]account.AccountResponse, error)
	GetYourRequest(ctx context.Context, from_id, to_id int64) ([]account.AccountResponse, error)
	GetYourFriend(ctx context.Context, from_id, to_id int64) ([]account.AccountResponse, error)
	UpdateStatus(ctx context.Context, user_id, from_id, to_id int64) (FollowResponse, error)
	UnFollow(ctx context.Context, from_id, to_id int64) error
}
