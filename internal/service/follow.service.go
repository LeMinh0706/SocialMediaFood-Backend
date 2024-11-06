package service

import (
	"context"
	"fmt"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/models"
)

type FollowService struct {
	queries        *db.Queries
	accountService *AccountService
}

func NewFollowService(queries *db.Queries, accountService *AccountService) (*FollowService, error) {
	return &FollowService{
		queries:        queries,
		accountService: accountService,
	}, nil
}

func (fs *FollowService) GetFollowStatus(ctx context.Context, from_id, to_id int64) (db.Follower, error) {
	var res db.Follower
	res, err := fs.queries.GetFollowStatus(ctx, db.GetFollowStatusParams{FromFollow: from_id, ToFollow: to_id})
	if err != nil {
		if err.Error() == "no rows in result set" {
			return db.Follower{FromFollow: from_id, ToFollow: to_id, Status: "No follow"}, nil
		}
		return res, err
	}
	return res, nil
}

func (fs *FollowService) CreateFollowRequest(ctx context.Context, from_id, to_id int64) (models.FollowRespone, error) {
	var res models.FollowRespone
	_, err := fs.accountService.GetAccountById(ctx, from_id)
	if err != nil {
		return res, err
	}
	_, err = fs.accountService.GetAccountById(ctx, to_id)
	if err != nil {
		return res, err
	}
	from, err := fs.queries.CreateFollow(ctx, db.CreateFollowParams{FromFollow: from_id, ToFollow: to_id, Status: "pending"})
	if err != nil {
		return res, err
	}
	to, err := fs.queries.CreateFollow(ctx, db.CreateFollowParams{FromFollow: to_id, ToFollow: from_id, Status: "accepted"})
	if err != nil {
		return res, err
	}
	res = models.FollowRespone{FromFollow: from, ToFollow: to}

	return res, nil
}

func (fs *FollowService) UpdateFriend(ctx context.Context, from_id, to_id int64) (models.FollowRespone, error) {
	var res models.FollowRespone
	req, err := fs.GetFollowStatus(ctx, from_id, to_id)
	if err != nil {
		return res, err
	}
	if req.Status == "friend" {
		return res, fmt.Errorf("their friend")
	}
	if req.Status != "accepted" {
		return res, fmt.Errorf("can't accept")
	}
	err = fs.queries.UpdateFriend(ctx, db.UpdateFriendParams{FromFollow: from_id, ToFollow: to_id})
	if err != nil {
		return res, err
	}
	from, err := fs.GetFollowStatus(ctx, from_id, to_id)
	if err != nil {
		return res, err
	}
	to, err := fs.GetFollowStatus(ctx, to_id, from_id)
	if err != nil {
		return res, err
	}
	res = models.FollowRespone{FromFollow: from, ToFollow: to}
	return res, nil

}

func (fs *FollowService) DeleteFollow(ctx context.Context, from_id, to_id int64) error {
	return fs.queries.DeleteFollow(ctx, db.DeleteFollowParams{FromFollow: from_id, ToFollow: to_id})
}
