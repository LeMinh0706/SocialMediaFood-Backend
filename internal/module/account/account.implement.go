package account

import (
	"context"
	"fmt"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type AccountService struct {
	queries *db.Queries
}

// SearchingAccount implements IAccountService.
func (a *AccountService) SearchingAccount(ctx context.Context, searching string, page int32, pageSize int32) ([]db.SearchingAccountsRow, error) {
	result, err := a.queries.SearchingAccounts(ctx, db.SearchingAccountsParams{
		Column1: pgtype.Text{String: searching, Valid: true},
		Limit:   pageSize,
		Offset:  (page - 1) * pageSize,
	})
	if err != nil {
		return []db.SearchingAccountsRow{}, err
	}
	return result, nil
}

// AddLocation implements IAccountService.
func (a *AccountService) AddLocation(ctx context.Context, user_id, account_id int64, lng string, lat string) (db.CreateOwnerBranchRow, error) {
	_, err := a.GetAccountAction(ctx, user_id, account_id)
	if err != nil {
		return db.CreateOwnerBranchRow{}, err
	}
	location := fmt.Sprintf("POINT(%s %s)", lat, lng)
	point, err := a.queries.CreateOwnerBranch(ctx, db.CreateOwnerBranchParams{
		AccountID:      account_id,
		StGeomfromtext: location,
	})
	if err != nil {
		return db.CreateOwnerBranchRow{}, err
	}
	return point, nil
}

// UpdateAvatar implements IAccountService.
func (a *AccountService) UpdateAvatar(ctx context.Context, id int64, user_id int64, url_avatar string) (AccountResponse, error) {
	var res AccountResponse
	_, err := a.GetAccountAction(ctx, id, user_id)
	if err != nil {
		return res, err
	}
	post, err := a.queries.CreatePost(ctx, db.CreatePostParams{
		PostTypeID: 3,
		AccountID:  id,
	})
	if err != nil {
		return res, err
	}
	image, err := a.queries.AddImagePost(ctx, db.AddImagePostParams{
		UrlImage: url_avatar,
		PostID:   post.ID,
	})
	if err != nil {
		return res, err
	}
	_, err = a.queries.UpdateAvatar(ctx, db.UpdateAvatarParams{
		ID:        id,
		UrlAvatar: image.UrlImage,
	})
	if err != nil {
		return res, err
	}
	account, _ := a.GetAccount(ctx, id)
	return account, nil
}

// UpdateBackground implements IAccountService.
func (a *AccountService) UpdateBackground(ctx context.Context, id int64, user_id int64, url_background string) (AccountResponse, error) {
	var res AccountResponse
	_, err := a.GetAccountAction(ctx, id, user_id)
	if err != nil {
		return res, err
	}
	post, err := a.queries.CreatePost(ctx, db.CreatePostParams{
		PostTypeID: 4,
		AccountID:  id,
	})
	if err != nil {
		return res, err
	}
	image, err := a.queries.AddImagePost(ctx, db.AddImagePostParams{
		UrlImage: url_background,
		PostID:   post.ID,
	})
	if err != nil {
		return res, err
	}
	_, err = a.queries.UpdateBackground(ctx, db.UpdateBackgroundParams{
		ID:                   id,
		UrlBackgroundProfile: image.UrlImage,
	})
	if err != nil {
		return res, err
	}
	account, _ := a.GetAccount(ctx, id)
	return account, nil
}

// UpdateName implements IAccountService.
func (a *AccountService) UpdateName(ctx context.Context, id int64, user_id int64, name string) (AccountResponse, error) {
	var res AccountResponse
	_, err := a.GetAccountAction(ctx, id, user_id)
	if err != nil {
		return res, err
	}
	_, err = a.queries.UpdateName(ctx, db.UpdateNameParams{
		ID:       id,
		Fullname: name,
	})
	if err != nil {
		return res, err
	}
	account, _ := a.GetAccount(ctx, id)
	return account, nil
}

// Backup implements IAccountService.
func (a *AccountService) Backup(ctx context.Context) {
	panic("unimplemented")
}

// GetAccount implements IAccountService.
func (a *AccountService) GetAccount(ctx context.Context, id int64) (AccountResponse, error) {
	acc, err := a.queries.GetDetailAccount(ctx, id)
	if err != nil {
		return AccountResponse{}, err
	}
	res := AccountRes(acc)
	return res, nil
}

// GetAccountAction implements IAccountService.
func (a *AccountService) GetAccountAction(ctx context.Context, id int64, user_id int64) (db.GetAccountByIdRow, error) {
	var res db.GetAccountByIdRow
	acc, err := a.queries.GetAccountById(ctx, id)
	if err != nil {
		return res, err
	}
	if acc.UserID != user_id {
		return res, fmt.Errorf("not you")
	}
	return acc, nil
}

// GetAccountById implements IAccountService.
func (a *AccountService) GetAccountById(ctx context.Context, id int64) (db.GetAccountByIdRow, error) {
	acc, err := a.queries.GetAccountById(ctx, id)
	if err != nil {
		return db.GetAccountByIdRow{}, err
	}
	return acc, nil
}

// GetAccountByUserId implements IAccountService.
func (a *AccountService) GetAccountByUserId(ctx context.Context, user_id int64) ([]AccountResponse, error) {
	res := make([]AccountResponse, 0)
	list, err := a.queries.GetAccountByUserId(ctx, user_id)
	if err != nil {
		return []AccountResponse{}, err
	}
	for _, element := range list {
		acc, _ := a.GetAccount(ctx, element)
		res = append(res, acc)
	}
	return res, nil
}

func NewAccountService(queries *db.Queries) IAccountService {
	return &AccountService{
		queries: queries,
	}
}
