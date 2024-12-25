package account

import (
	"time"

	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type CreateAccountVip struct {
	UserID               int64  `json:"user_id"`
	Fullname             string `json:"fullname"`
	Country              string `json:"country"`
	Language             string `json:"language"`
	UrlAvatar            string `json:"url_avatar"`
	UrlBackgroundProfile string `json:"url_background_profile"`
}

type EmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type AccountResponse struct {
	ID                   int64  `json:"id"`
	Fullname             string `json:"fullname"`
	UrlAvatar            string `json:"url_avatar"`
	UrlBackgroundProfile string `json:"url_background_profile"`
	Gender               int32  `json:"gender"`
	Country              string `json:"country"`
	Language             string `json:"language"`
	Address              string `json:"address"`
	RoleID               int32  `json:"role_id"`
	IsUpgrade            bool   `json:"is_upgrade"`
}

type GetMeResponse struct {
	Accounts    []AccountResponse `json:"accounts"`
	Email       string            `json:"email"`
	AccessToken string            `json:"access_token"`
}

type UpdateNameReq struct {
	Fullname string `json:"fullname"`
}

type UpgradeOwnerRequest struct {
	AccountID int64 `json:"account_id"`
}

type UpgradePrice struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Benefit   string    `json:"benefit" binding:"required"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func AccountRes(account db.Account) AccountResponse {
	return AccountResponse{
		ID:                   account.ID,
		Fullname:             account.Fullname,
		UrlAvatar:            account.UrlAvatar,
		UrlBackgroundProfile: account.UrlBackgroundProfile,
		Gender:               account.Gender.Int32,
		Country:              account.Country.String,
		Language:             account.Language.String,
		Address:              account.Address.String,
		RoleID:               account.RoleID,
		IsUpgrade:            account.IsUpgrade.Bool,
	}
}

func ListAccountResponse(all []db.Account) []AccountResponse {
	var list []AccountResponse
	for _, acc := range all {
		res := AccountRes(acc)
		list = append(list, res)
	}
	return list
}
