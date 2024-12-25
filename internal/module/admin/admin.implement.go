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

// ChoosingPrice implements IAdminService.
func (a *AdminService) ChoosingPrice(ctx context.Context, username string, id int64) error {
	err := a.IsAdmin(ctx, username)
	if err != nil {
		return err
	}
	err = a.queries.PriceChoosing(ctx, id)
	if err != nil {
		return err
	}
	err = a.queries.UnableChoose(ctx)
	if err != nil {
		return err
	}
	return nil
}

// GetDetailReportPost implements IAdminService.
func (a *AdminService) GetDetailReportPost(ctx context.Context, username string, post_id int64, account_id int64, page int32, page_size int32) (ReportDetailResponse, error) {
	var res ReportDetailResponse
	err := a.IsAdmin(ctx, username)
	if err != nil {
		return res, err
	}
	post, err := a.post.GetPost(ctx, account_id, post_id)
	if err != nil {
		return res, err
	}
	list, err := a.queries.ReportPostDetail(ctx, db.ReportPostDetailParams{
		PostID: post_id,
		Limit:  page_size,
		Offset: (page - 1) * page_size,
	})
	if err != nil {
		return res, err
	}
	arr := make([]ReportFrom, 0)
	for _, element := range list {
		acc, _ := a.acc.GetAccountById(ctx, element.AccountID)
		i := ReportFrom{Account: acc, Issue: db.IssuePost{ID: element.IssueID, Name: element.Name.String, IsDeleted: element.IsDeleted.Bool}}
		arr = append(arr, i)
	}
	res = ReportDetailResponse{Post: post, IssuePosts: arr}
	return res, nil
}

// GetUpgradeSuccess implements IAdminService.
func (a *AdminService) GetUpgradeSuccess(ctx context.Context, page int32, page_size int32) ([]account.AccountResponse, error) {
	panic("unimplemented")
}

// AddUpgragePrice implements IAdminService.
func (a *AdminService) AddUpgragePrice(ctx context.Context, username string, title string, benefit string, price float64) (UpgradePrice, error) {
	err := a.IsAdmin(ctx, username)
	if err != nil {
		return UpgradePrice{}, err
	}
	create, err := a.queries.AddUpgradePrice(ctx, db.AddUpgradePriceParams{
		Title:   title,
		Price:   pgtype.Numeric{Exp: -3, Int: big.NewInt(int64(price * 1000)), Valid: true},
		Benefit: benefit,
	})
	if err != nil {
		return UpgradePrice{}, err
	}
	priceFloat, _ := create.Price.Float64Value()
	return UpgradePrice{ID: create.ID, Price: priceFloat.Float64, CreatedAt: create.CreatedAt.Time}, nil
}

// BanPost implements IAdminService.
func (a *AdminService) BanPost(ctx context.Context, username string, post_id int64) (post.PostResponse, error) {
	panic("unimplemented")
}

// GetListReportPost implements IAdminService.
func (a *AdminService) GetListReportPost(ctx context.Context, username string, account_id int64, page int32, page_size int32) ([]post.PostResponse, error) {
	var res []post.PostResponse
	err := a.IsAdmin(ctx, username)
	if err != nil {
		return nil, err
	}
	list, err := a.queries.GetListPostReport(ctx, db.GetListPostReportParams{
		Limit:  page_size,
		Offset: (page - 1) * page_size,
	})
	if err != nil {
		return nil, err
	}
	for _, element := range list {
		post, err := a.post.GetPost(ctx, account_id, element.PostID)
		if err != nil {
			return nil, err
		}
		res = append(res, post)
	}
	return res, nil
}

// GetUpgradePrice implements IAdminService.
func (a *AdminService) GetUpgradePrice(ctx context.Context, page int32, page_size int32) ([]UpgradePrice, error) {
	var res []UpgradePrice
	list, err := a.queries.GetListUpgradePrice(ctx, db.GetListUpgradePriceParams{
		Limit:  page_size,
		Offset: (page - 1) * page_size,
	})
	if err != nil {
		return nil, err
	}
	for _, element := range list {
		priceFloat, _ := element.Price.Float64Value()
		add := UpgradePrice{ID: element.ID, Price: priceFloat.Float64, CreatedAt: element.CreatedAt.Time, Title: element.Title, Benefit: element.Benefit, IsChoose: element.IsChoose}
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
func (a *AdminService) IsAdmin(ctx context.Context, username string) error {
	user, err := a.queries.Login(ctx, username)
	if err != nil {
		return fmt.Errorf("permission")
	}
	role, err := a.queries.IsAdmin(ctx, user.ID)
	if err != nil {
		return err
	}
	if role != 1 {
		return fmt.Errorf("permission")
	}
	return nil
}

// RejectBan implements IAdminService.
func (a *AdminService) RejectBan(ctx context.Context, username string, post_id int64) error {
	panic("unimplemented")
}

// UpgradeReject implements IAdminService.
func (a *AdminService) UpgradeReject(ctx context.Context, username string, account_id int64) error {
	err := a.IsAdmin(ctx, username)
	if err != nil {
		return err
	}
	err = a.queries.UpgradeReject(ctx, account_id)
	if err != nil {
		return err
	}
	return nil
}

// UpgradeSuccess implements IAdminService.
func (a *AdminService) UpgradeSuccess(ctx context.Context, username string, account_id int64) (account.AccountResponse, error) {
	var res account.AccountResponse
	err := a.IsAdmin(ctx, username)
	if err != nil {
		return res, err
	}
	_, err = a.queries.GetStatusQueue(ctx, account_id)
	if err != nil {
		return res, fmt.Errorf("they're not request to upgrade")
	}
	err = a.queries.UpgradeStateQueue(ctx, account_id)
	if err != nil {
		return res, err
	}
	err = a.queries.UpgradeOwner(ctx, account_id)
	if err != nil {
		return res, err
	}
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
