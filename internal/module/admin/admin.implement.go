package admin

import (
	"context"
	"fmt"
	"math/big"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/account"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/post"
	"github.com/jackc/pgx/v5/pgtype"
)

type AdminService struct {
	queries *db.Queries
	acc     account.IAccountService
	post    post.IPostService
}

// AddUpgragePrice implements IAdminService.
func (a *AdminService) AddUpgragePrice(ctx context.Context, user_id int64, price float64) (UpgradePrice, error) {
	err := a.IsAdmin(ctx, user_id)
	if err != nil {
		return UpgradePrice{}, fmt.Errorf("not admin")
	}
	create, err := a.queries.AddUpgradePrice(ctx, pgtype.Numeric{Exp: -3, Int: big.NewInt(int64(price * 1000))})
	if err != nil {
		return UpgradePrice{}, err
	}
	priceFloat, _ := create.Price.Float64Value()
	return UpgradePrice{ID: create.ID, Price: priceFloat.Float64, CreatedAt: create.CreatedAt.Time}, nil
}

// BanPost implements IAdminService.
func (a *AdminService) BanPost(ctx context.Context, user_id int64, post_id int64) (post.PostResponse, error) {
	panic("unimplemented")
}

// GetListReportPost implements IAdminService.
func (a *AdminService) GetListReportPost(ctx context.Context, user_id int64, account_id int64, page int32, page_size int32) ([]post.PostResponse, error) {
	panic("unimplemented")
}

// GetUpgradePrice implements IAdminService.
func (a *AdminService) GetUpgradePrice(ctx context.Context, page int32, page_size int32) ([]UpgradePrice, error) {
	var res []UpgradePrice
	list, err := a.queries.GetUpgradePrice(ctx, db.GetUpgradePriceParams{
		Limit:  page_size,
		Offset: (page - 1) * page_size,
	})
	if err != nil {
		return nil, err
	}
	for _, element := range list {
		priceFloat, _ := element.Price.Float64Value()
		add := UpgradePrice{ID: element.ID, Price: priceFloat.Float64, CreatedAt: element.CreatedAt.Time}
		res = append(res, add)
	}
	return res, err
}

// GetUpgradeQueue implements IAdminService.
func (a *AdminService) GetUpgradeQueue(ctx context.Context, page int32, page_size int32) ([]account.AccountResponse, error) {
	var res []account.AccountResponse
	list, err := a.queries.GetUpgradeQueue(ctx, db.GetUpgradeQueueParams{
		Limit:  page_size,
		Offset: (page - 1) * page_size,
	})
	if err != nil {
		return nil, err
	}
	for _, element := range list {
		acc, _ := a.acc.GetAccount(ctx, element)
		res = append(res, acc)
	}
	return res, nil
}

// IsAdmin implements IAdminService.
func (a *AdminService) IsAdmin(ctx context.Context, user_id int64) error {
	role, err := a.queries.IsAdmin(ctx, user_id)
	if err != nil {
		return err
	}
	if role != 1 {
		return fmt.Errorf("permission")
	}
	return nil
}

// RejectBan implements IAdminService.
func (a *AdminService) RejectBan(ctx context.Context, user_id int64, post_id int64) error {
	panic("unimplemented")
}

// UpgradeReject implements IAdminService.
func (a *AdminService) UpgradeReject(ctx context.Context, user_id int64, account_id int64) (account.AccountResponse, error) {
	panic("unimplemented")
}

// UpgradeSuccess implements IAdminService.
func (a *AdminService) UpgradeSuccess(ctx context.Context, user_id int64, account_id int64) (account.AccountResponse, error) {
	var res account.AccountResponse
	err := a.IsAdmin(ctx, user_id)
	if err != nil {
		return res, err
	}
	go a.queries.UpgradeStateQueue(ctx, account_id)
	go a.queries.UpgradeOwner(ctx, account_id)
	account, _ := a.acc.GetAccount(ctx, account_id)
	return account, nil
}

func NewAdminService(q *db.Queries, a account.IAccountService, p post.IPostService) IAdminService {
	return &AdminService{
		queries: q,
		acc:     a,
		post:    p,
	}
}
