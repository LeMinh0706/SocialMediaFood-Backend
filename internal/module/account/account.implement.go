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

// UpgradeOwnerRequest implements IAccountService.
func (a *AccountService) UpgradeOwnerRequest(ctx context.Context, id int64, username string) error {
	acc, err := a.GetAccountAction(ctx, id, username)
	if acc.RoleID != 3 {
		return fmt.Errorf("your account can't upgrade")
	}
	if err != nil {
		return err
	}
	price_id, err := a.queries.GetChoosePrice(ctx)
	if err != nil {
		return err
	}
	err = a.queries.UpgradeOnwerRequest(ctx, db.UpgradeOnwerRequestParams{
		AccountID:      id,
		UpgradePriceID: price_id.ID,
	})
	if err != nil {
		return err
	}
	return nil
}

// UpdateEmail implements IAccountService.
func (a *AccountService) UpdateEmail(ctx context.Context, id int64, username string, email string) error {
	_, err := a.GetAccountAction(ctx, id, username)
	if err != nil {
		return err
	}
	err = a.queries.UpdateEmail(ctx, db.UpdateEmailParams{
		ID:    id,
		Email: pgtype.Text{String: email, Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

// AddLocation implements IAccountService.
func (a *AccountService) AddLocation(ctx context.Context, account_id int64, username string, address string, lng string, lat string) (db.CreateOwnerBranchRow, error) {
	_, err := a.GetAccountAction(ctx, account_id, username)
	if err != nil {
		return db.CreateOwnerBranchRow{}, err
	}
	location := fmt.Sprintf("POINT(%s %s)", lat, lng)
	point, err := a.queries.CreateOwnerBranch(ctx, db.CreateOwnerBranchParams{
		AccountID:      account_id,
		StGeomfromtext: location,
		Address:        address,
	})
	if err != nil {
		return db.CreateOwnerBranchRow{}, err
	}
	return point, nil
}

// UpdateBackground implements IAccountService.
func (a *AccountService) UpdateBackground(ctx context.Context, id int64, username string, url_background string) (AccountResponse, error) {
	var res AccountResponse
	_, err := a.GetAccountAction(ctx, id, username)
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

// UpdateAvatar implements IAccountService.
func (a *AccountService) UpdateAvatar(ctx context.Context, id int64, username string, url_avatar string) (AccountResponse, error) {
	var res AccountResponse
	_, err := a.GetAccountAction(ctx, id, username)
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

// UpdateName implements IAccountService.
func (a *AccountService) UpdateName(ctx context.Context, id int64, username string, name string) (AccountResponse, error) {
	var res AccountResponse

	_, err := a.GetAccountAction(ctx, id, username)
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

// GetAccountAction implements IAccountService.
func (a *AccountService) GetAccountAction(ctx context.Context, id int64, username string) (db.Account, error) {
	var res db.Account
	user, err := a.queries.Login(ctx, username)
	if err != nil {
		return res, fmt.Errorf("not you")
	}
	acc, err := a.queries.GetDetailAccount(ctx, id)
	if err != nil {
		return res, err
	}
	if acc.UserID != user.ID {
		return res, fmt.Errorf("not you")
	}
	return acc, nil
}

// GetAccountByUserId implements IAccountService.
func (a *AccountService) GetAccountByUserId(ctx context.Context, username string) (GetMeResponse, error) {
	var res GetMeResponse
	res.Accounts = make([]AccountResponse, 0)
	user, err := a.queries.Login(ctx, username)
	if err != nil {
		return res, fmt.Errorf("not you")
	}
	list, err := a.queries.GetAccountByUserId(ctx, user.ID)
	if err != nil {
		return res, err
	}
	for _, element := range list {
		account, _ := a.queries.GetDetailAccount(ctx, element)
		add := AccountRes(account)
		res.Accounts = append(res.Accounts, add)
	}

	res.Email = user.Email.String
	return res, nil
}

// AddEmail implements IAccountService.
func (a *AccountService) AddEmail(ctx context.Context, username string, email string) error {
	user, err := a.queries.Login(ctx, username)
	if err != nil {
		return fmt.Errorf("not you")
	}
	err = a.queries.AddEmail(ctx, db.AddEmailParams{
		ID:    user.ID,
		Email: pgtype.Text{String: email, Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
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

// GetAccount implements IAccountService.
func (a *AccountService) GetAccount(ctx context.Context, id int64) (AccountResponse, error) {
	acc, err := a.queries.GetDetailAccount(ctx, id)
	if err != nil {
		return AccountResponse{}, err
	}
	res := AccountRes(acc)
	return res, nil
}

// GetAccountById implements IAccountService.
func (a *AccountService) GetAccountById(ctx context.Context, id int64) (db.GetAccountByIdRow, error) {
	acc, err := a.queries.GetAccountById(ctx, id)
	if err != nil {
		return db.GetAccountByIdRow{}, err
	}
	return acc, nil
}

func (a *AccountService) GetUpgradePrice(ctx context.Context) (db.GetChoosePriceRow, error) {
	res, err := a.queries.GetChoosePrice(ctx)
	if err != nil {
		return db.GetChoosePriceRow{}, err
	}
	return res, nil
}

func NewAccountService(queries *db.Queries) IAccountService {
	return &AccountService{
		queries: queries,
	}
}
