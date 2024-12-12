package follower

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type IFollowerService interface {
	Backup(ctx context.Context)
	FollowRequest(ctx context.Context, username string, from_id, to_id int64) (FollowResponse, error)
	UtilCreateFollow(ctx context.Context, from, to db.Follower) (AccountFollowResponse, AccountFollowResponse, error)
	GetRequestStatus(ctx context.Context, arg db.GetFollowStatusParams) (db.GetFollowStatusRow, error)
	GetYourFollower(ctx context.Context, page, page_size int32, from_id int64) (ListFollow, error)
	GetYourRequest(ctx context.Context, page, page_size int32, from_id int64) (ListFollow, error)
	GetYourFriend(ctx context.Context, page, page_size int32, from_id int64) (ListFollow, error)
	GetFollowType(ctx context.Context, status string, page, page_size int32, from_id int64) (ListFollow, error)
	UpdateStatus(ctx context.Context, username string, from_id, to_id int64) error
	UnFollow(ctx context.Context, username string, from_id, to_id int64) error
}
