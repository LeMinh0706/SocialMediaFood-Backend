package account

import (
	"github.com/LeMinh0706/SocialMediaFood-Backend/db"
)

type EmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type AccountResponse struct {
	ID                   int64  `json:"id"`
	UserID               int64  `json:"user_id"`
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
	Accounts []AccountResponse `json:"accounts"`
	Email    string            `json:"email"`
}

type UpdateNameReq struct {
	Fullname string `json:"fullname"`
}

type UpgradeOwnerRequest struct {
	AccountID int64 `json:"account_id"`
}

func AccountRes(account db.Account) AccountResponse {
	return AccountResponse{
		ID:                   account.ID,
		UserID:               account.UserID,
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
