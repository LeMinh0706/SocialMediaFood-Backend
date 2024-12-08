package admin

import (
	"context"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/account"
	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/module/post"
)

type IAdminService interface {
	IsAdmin(ctx context.Context, user_id int64) error
	GetUpgradePrice(ctx context.Context, page, page_size int32) ([]UpgradePrice, error)
	AddUpgragePrice(ctx context.Context, user_id int64, title, benefit string, price float64) (UpgradePrice, error)
	GetUpgradeQueue(ctx context.Context, page, page_size int32) ([]account.AccountResponse, error)
	GetUpgradeSuccess(ctx context.Context, page, page_size int32) ([]account.AccountResponse, error)
	UpgradeSuccess(ctx context.Context, user_id, account_id int64) (account.AccountResponse, error)
	UpgradeReject(ctx context.Context, user_id, account_id int64) (account.AccountResponse, error)
	GetListReportPost(ctx context.Context, user_id, account_id int64, page, page_size int32) ([]post.PostResponse, error)
	GetDetailReportPost(ctx context.Context, post_id, user_id, account_id int64, page, page_size int32) (ReportDetailResponse, error)
	ChoosingPrice(ctx context.Context, user_id, id int64) error
	BanPost(ctx context.Context, user_id, post_id int64) (post.PostResponse, error)
	RejectBan(ctx context.Context, user_id, post_id int64) error
}
