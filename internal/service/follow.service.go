package service

import (
	"context"
	"fmt"
	"strconv"

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

func (fs *FollowService) GetFollow(ctx context.Context, from_id, pageStr, pageSizeStr string) (models.ListFollow, error) {
	var res models.ListFollow
	id, err := strconv.ParseInt(from_id, 10, 64)
	if err != nil {
		return res, fmt.Errorf("id number")
	}
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		return res, fmt.Errorf("page number")
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		return res, fmt.Errorf("pagesize number")
	}
	list, err := fs.queries.GetYourFollow(ctx, db.GetYourFollowParams{FromFollow: id, Limit: int32(pageSize), Offset: (int32(page) - 1) * int32(pageSize)})
	if err != nil {
		return res, err
	}
	accRes := make([]models.AccountForPost, 0)
	for _, element := range list {
		acc, err := fs.accountService.GetAccountById(ctx, element)
		if err != nil {
			return res, err
		}
		accRes = append(accRes, acc)
	}
	total, err := fs.queries.CountFollow(ctx, id)
	if err != nil {
		return res, err
	}
	res = models.ListFollow{
		YourFollows: accRes,
		Total:       total,
	}
	return res, nil
}

func (fs *FollowService) GetFollower(ctx context.Context, from_id, pageStr, pageSizeStr string) (models.ListFollow, error) {
	var res models.ListFollow
	id, err := strconv.ParseInt(from_id, 10, 64)
	if err != nil {
		return res, fmt.Errorf("id number")
	}
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		return res, fmt.Errorf("page number")
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		return res, fmt.Errorf("pagesize number")
	}
	list, err := fs.queries.GetYourFollower(ctx, db.GetYourFollowerParams{FromFollow: id, Limit: int32(pageSize), Offset: (int32(page) - 1) * int32(pageSize)})
	if err != nil {
		return res, err
	}
	accRes := make([]models.AccountForPost, 0)
	for _, element := range list {
		acc, err := fs.accountService.GetAccountById(ctx, element)
		if err != nil {
			return res, err
		}
		accRes = append(accRes, acc)
	}
	total, err := fs.queries.CountFollower(ctx, id)
	if err != nil {
		return res, err
	}
	res = models.ListFollow{
		YourFollows: accRes,
		Total:       total,
	}
	return res, nil
}

func (fs *FollowService) GetFriend(ctx context.Context, from_id, pageStr, pageSizeStr string) (models.ListFollow, error) {
	var res models.ListFollow
	id, err := strconv.ParseInt(from_id, 10, 64)
	if err != nil {
		return res, fmt.Errorf("id number")
	}
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		return res, fmt.Errorf("page number")
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		return res, fmt.Errorf("pagesize number")
	}
	list, err := fs.queries.GetYourFriend(ctx, db.GetYourFriendParams{FromFollow: id, Limit: int32(pageSize), Offset: (int32(page) - 1) * int32(pageSize)})
	if err != nil {
		return res, err
	}
	accRes := make([]models.AccountForPost, 0)
	for _, element := range list {
		acc, err := fs.accountService.GetAccountById(ctx, element)
		if err != nil {
			return res, err
		}
		accRes = append(accRes, acc)
	}
	total, err := fs.queries.CountFriend(ctx, id)
	if err != nil {
		return res, err
	}
	res = models.ListFollow{
		YourFollows: accRes,
		Total:       total,
	}
	return res, nil
}

func (fs *FollowService) CreateFollowRequest(ctx context.Context, from_id, to_id int64) (models.FollowRespone, error) {
	var res models.FollowRespone
	check, err := fs.GetFollowStatus(ctx, from_id, to_id)
	if check.Status == "accepted" {
		return res, fmt.Errorf("waiting accept")
	}
	if err != nil {
		return res, nil
	}

	_, err = fs.accountService.GetAccountById(ctx, from_id)
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
